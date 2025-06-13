package main

import (
	"api/database"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

func PostLogin(w http.ResponseWriter, r *http.Request) {
	// Decode login
	var loginForm UserCredsForm
	err := json.NewDecoder(r.Body).Decode(&loginForm)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate login credentials with database
	filter := bson.M{"Email": loginForm.Email, "Password": loginForm.Password}
	singleResult := database.GetCollectionFromMongo(database.Users).FindOne(r.Context(), filter)
	var message string
	var foundUser UserCredsForm
	err = singleResult.Decode(&foundUser)
	if err != nil {
		http.Error(w, "User not found", http.StatusForbidden)
		return
	} else {
		message = fmt.Sprintf("Login successful. Welcome %s", foundUser.Email)
	}

	// foundUser has a valid user, create token
	token, err := createUserToken(&foundUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Create user token failed: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// w.WriteHeader(http.StatusOK) // 200
	json.NewEncoder(w).Encode(map[string]string{
		"message":       message,
		"authenticated": strconv.FormatBool(true),
		"token":         token,
	})
}
