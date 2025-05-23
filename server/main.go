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
	router.HandleFunc("/login", PostLogin) // localhost:8000/login POST
	router.HandleFunc("/signup", PostSignUp)

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
