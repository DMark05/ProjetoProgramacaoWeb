package main

import (
	"api/database"
	"encoding/json"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginForm struct {
	Email    string `json:"Email" bson:"Email"`
	Password string `json:"Password" bson:"Password"`
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var loginForm LoginForm

	err := json.NewDecoder(r.Body).Decode(&loginForm)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	
	filter := bson.M{"Email": loginForm.Email, "Password": loginForm.Password}
	singleResult := database.GetCollectionFromMongo(database.Users).FindOne(r.Context(), filter)

	var foundUser LoginForm
	err = singleResult.Decode(&foundUser)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Login successful. Welcome %s", foundUser.Email)
}
