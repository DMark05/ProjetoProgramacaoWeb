package main

import (
	"io"
	"strings"
	"net/http"
	"encoding/json"
)

type LoginForm struct {
	Email string `bson:"e-mail"`
	Password string `bson:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var login_form LoginForm = LoginForm{}

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

	_ , err = getCollectionFromMongo(Users).InsertOne(r.Context(), login_form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}