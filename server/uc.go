package main

import (
	"api/database"
	"encoding/json"
	"net/http"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type UCStruct struct {
	Name	string `json:"Name" bson:"Name"`
	Year	int `json:"Year" bson:"Year"` // Cadeira 1º, 2º ou 3º ano
	Description string `json:"Description" bson:"Description"`
	Teacher string `json:"Teacher" bson:"Teacher"`
	Image	string `json:"Image" bson:"Image"`
}

func PostCreateUC (w http.ResponseWriter, r *http.Request) {
	var newUC UCStruct
	err := json.NewDecoder(r.Body).Decode(&newUC)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	filter := bson.M{
					"Name": newUC.Name,
					"Year": newUC.Year,
					"Description": newUC.Description,
					"Teacher": newUC.Teacher,
					"Image": newUC.Image}

	_, insertErr := database.GetCollectionFromMongo(database.UC).InsertOne(r.Context(), filter)
	if insertErr != nil {
		http.Error(w, "Creating UC failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("UC %s created", newUC.Name),
	})
}