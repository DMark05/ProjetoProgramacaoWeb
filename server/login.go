package main

import (
	"api/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type LoginForm struct {
	Email    string `bson:"e-mail"`
	Password string `bson:"password"`
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	login_form := LoginForm{}
	output, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(output, &login_form)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}
	singleResult := database.GetCollectionFromMongo(database.Users).FindOne(r.Context(), login_form)
	login_form = LoginForm{}
	err = singleResult.Decode(&login_form)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding result from mongoDB: %v", err), http.StatusInternalServerError)
	}
	if login_form.Email == "" {
		http.Error(w, "Wrong Credentials", http.StatusBadRequest)
	}
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	login_form := LoginForm{}
	output, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	err = json.Unmarshal(output, &login_form)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
	}
	domain := strings.Split(login_form.Email, "@")[1]
	if domain != "iscte-iul.pt" {
		http.Error(w, "Invalid data", http.StatusInternalServerError)
	}
	result, err:= database.GetCollectionFromMongo(database.Users).InsertOne(r.Context(), login_form)
	// Print debug
	if err != nil {
		fmt.Printf("Inserted document with _id %v\n", result.InsertedID)
	}
}
