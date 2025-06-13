package main

import (
	"api/database"
	"api/middleware"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatalln(err)
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /login", PostLogin) // localhost:8000/login POST
	router.HandleFunc("POST	/signup", PostSignUp)
	router.Handle("GET /events", middleware.ValidateJWT(http.HandlerFunc(GetEvent), verifyUserToken))
	router.Handle("POST /newevent", middleware.ValidateJWT(http.HandlerFunc(PostcreateEvent), verifyUserToken))
	router.Handle("GET /events/{eventId}", middleware.ValidateJWT(http.HandlerFunc(getEventImage), verifyUserToken))
	router.Handle("GET /events/{eventId}/reviews", middleware.ValidateJWT(http.HandlerFunc(getEventReviews), verifyUserToken))
	router.Handle("PUT /events/{eventId}/reviews", middleware.ValidateJWT(http.HandlerFunc(PutEventReview), verifyUserToken))
	// router.HandleFunc("/events", GetEvent)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	finalHandler := corsHandler.Handler(middleware.RedirectDefaultWrongCallsMiddleware(router))

	server := http.Server{
		Handler: finalHandler,
		Addr:    ":5000",
	}
	if err != server.ListenAndServe() {
		log.Fatalln(err)
	}
}
