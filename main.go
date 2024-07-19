package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chandanjnv/go-movie-crud/handlers"
	"github.com/chandanjnv/go-movie-crud/models"
	"github.com/gorilla/mux"
)

func init() {
	models.InitializeMovies()
}

// main function initializes the server and routes.
func main() {
	port := "8080"
	r := mux.NewRouter()

	// Define HTTP routes and handlers.
	r.HandleFunc("/movies", handlers.GetMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", handlers.CreateMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods(http.MethodDelete)

	// Start the HTTP server.
	fmt.Printf("Starting Server at port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
