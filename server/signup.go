package main

import (
	"api/database"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type SignUpForm struct {
	Email    string `json:"Email" bson:"Email"`
	Password string `json:"Password" bson:"Password"`
}

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	var signUpForm SignUpForm

	err := json.NewDecoder(r.Body).Decode(&signUpForm)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	filter := bson.M{"Email": signUpForm.Email, "Password": signUpForm.Password}
	singleResult := database.GetCollectionFromMongo(database.Users).FindOne(r.Context(), filter)

	var user SignUpForm

	err = singleResult.Decode(&user)
	if errors.Is(err, mongo.ErrNoDocuments) {
		_, insertErr := database.GetCollectionFromMongo(database.Users).InsertOne(r.Context(), signUpForm)
		if insertErr != nil {
			http.Error(w, "Sign up failed", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("Sign up successful. Welcome %s", signUpForm.Email),
		})
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	} else if signUpForm.Email == user.Email {
		http.Error(w, "Email in use", http.StatusConflict)
		return
	}
}
