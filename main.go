package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"MyService/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/api/v1/movies", handlers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/v1/movie/{id}", handlers.GetMovie).Methods("GET")
	router.HandleFunc("/api/v1/movie", handlers.AddMovie()).Methods("POST")
	router.HandleFunc("/api/v1/movies/batch", handlers.AddMoviesInBatch()).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}", handlers.UpdateMovie()).Methods("PUT")
	router.HandleFunc("/api/v1/movies/batch", handlers.UpdateMoviesInBatch()).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}", handlers.DeleteMovie()).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/batch", handlers.DeleteMoviesInBatch()).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/search", handlers.SearchMovies()).Queries("t", "{t}").Queries("y", "{y}").Queries("d", "{d}").Queries("i", "{i}").Queries("g", "{g}").Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
