package main

import (
	"api/database"
	"api/middleware"
	"log"
	"net/http"
)

func main() {
	err := database.Init()
	if err != nil {
		log.Fatalln(err)
	}
	router := http.NewServeMux()
	router.HandleFunc("/login", PostLogin) // localhost:8000/login POST
	server := http.Server{
		Handler: middleware.RedirectDefaultWrongCallsMiddleware(router),
		Addr:    ":5000",
	}
	if err != server.ListenAndServe() {
		log.Fatalln(err)
	}
}
