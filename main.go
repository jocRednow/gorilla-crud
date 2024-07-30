package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json"id"`
	Isbn     string    `json"isbn"`
	Title    string    `json"title"`
	Director *Director `json"director"`
}

type Director struct {
	Firstname string `json"firstname"`
	Lastname  string `json"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Text", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Text", "application/json")
	params := mux.Vars(r)
	for _, v := range movies {
		if v.ID == params["id"] {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Text", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Text", "application/json")
	params := mux.Vars(r)
	for i, v := range movies {
		if v.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Text", "application/json")
	params := mux.Vars(r)
	for i, v := range movies {
		if v.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			// movies = append(movies[:i], movies...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1200", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
