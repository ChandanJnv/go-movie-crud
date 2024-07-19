package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/chandanjnv/go-movie-crud/models"
	"github.com/gorilla/mux"
)

// GetMovies handles the HTTP GET request to retrieve all movies.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.MovieList)
}

// GetMovie handles the HTTP GET request to retrieve a specific movie by ID.
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range models.MovieList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreateMovie handles the HTTP POST request to add a new movie.
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000)) // Generate a random ID for the new movie.
	models.MovieList = append(models.MovieList, movie)
	json.NewEncoder(w).Encode(movie)
}

// UpdateMovie handles the HTTP PUT request to update an existing movie by ID.
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.MovieList {
		if item.ID == params["id"] {
			models.MovieList = append(models.MovieList[:index], models.MovieList[index+1:]...) // Remove the existing movie.
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"] // Preserve the existing movie ID.
			models.MovieList = append(models.MovieList, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// DeleteMovie handles the HTTP DELETE request to remove a movie by ID.
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.MovieList {
		if item.ID == params["id"] {
			models.MovieList = append(models.MovieList[:index], models.MovieList[index+1:]...) // Remove the movie.
			break
		}
	}
	json.NewEncoder(w).Encode(models.MovieList)
}
