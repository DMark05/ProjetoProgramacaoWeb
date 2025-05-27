package main

import (
	"api/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EventStruct struct {
	Name	string `json:"Name" bson:"Name"`
	Date	time.Time `json:"Date" bson:"Date"`
	Description string `json:"Description" bson:"Description"`
	Organizer string `json:"Organizer" bson:"Organizer"`
	Tags	[]string `json:"Tags" bson:"Tags"`
	Image	string `json:"Image" bson:"Image"`
}

func PostcreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent EventStruct
	err := json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	filter := bson.M{
					"Name": newEvent.Name,
					"Date": newEvent.Date,
					"Description": newEvent.Description,
					"Organizer": newEvent.Organizer,
					"Tags": newEvent.Tags,
					"Image": newEvent.Image}

	_, insertErr := database.GetCollectionFromMongo(database.Events).InsertOne(r.Context(), filter)
	if insertErr != nil {
			http.Error(w, "Creating event failed", http.StatusInternalServerError)
			return
		}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Event %s created", newEvent.Name),
	})
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	// Default values
	page := 1
	limit := 10
	
	var err error

	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}
	}

	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil || limit < 1 {
			limit = 10
		}
	}

	skip := (page - 1) * limit

	cursor, err := database.GetCollectionFromMongo(database.Events).Find(
		r.Context(),
		bson.M{},
		options.Find().
			SetSkip(int64(skip)).
			SetLimit(int64(limit)),
	)
	if err != nil {
		http.Error(w, "Error geting events", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	var events []EventStruct
	if err := cursor.All(r.Context(), &events); err != nil {
		http.Error(w, "Error processing events", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}