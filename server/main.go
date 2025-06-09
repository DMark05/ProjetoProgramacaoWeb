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
	router.HandleFunc("/login", PostLogin)
	router.HandleFunc("/signup", PostSignUp)
	router.Handle("/events", middleware.ValidateJWT(http.HandlerFunc(GetEvent), verifyUserToken))
	router.Handle("/newevent", middleware.ValidateJWT(http.HandlerFunc(PostcreateEvent), verifyUserToken))
	router.Handle("/events/{eventId}", middleware.ValidateJWT(http.HandlerFunc(getEventImage), verifyUserToken))
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
