package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"MyService/models"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to My Service, Author: Ankit Kumar!"))
	if err != nil {
		log.Println("Error in writing to response writer")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	if len(models.MapOfMovies) == 0 {
		log.Println("No movies found")
		http.Error(w, "No movies found", http.StatusNotFound)
		return
	}
	log.Println("retrieving all movies")
	err := json.NewEncoder(w).Encode(models.MapOfMovies)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Println("retrieving movie with id: " + id)
	found := false
	for _, movie := range models.MapOfMovies {
		if _, ok := movie[id]; ok {
			found = true
			err := json.NewEncoder(w).Encode(movie[id])
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	if !found {
		log.Println("Movie not found with id: " + id)
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
}

func AddMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		log.Println("adding movie with body: " + string(reqBody))
		var movie models.Movie
		err := json.Unmarshal(reqBody, &movie)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movie.ID = uuid.New()
		models.MapOfMovies = append(models.MapOfMovies, map[string]models.Movie{movie.ID.String(): movie})
		_, err = w.Write([]byte("movie added"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func AddMoviesInBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		log.Println("adding movies in batch with body: " + string(reqBody))
		var movies []models.Movie
		err := json.Unmarshal(reqBody, &movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, movie := range movies {
			movie.ID = uuid.New()
			models.MapOfMovies = append(models.MapOfMovies, map[string]models.Movie{movie.ID.String(): movie})
		}
		_, err = w.Write([]byte("movies added"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		log.Println("updating movie with id: " + id)
		found := false
		for i, movieMap := range models.MapOfMovies {
			if _, ok := movieMap[id]; ok {
				found = true
				reqBody, _ := ioutil.ReadAll(r.Body)
				log.Println("updating movie with body: " + string(reqBody))
				var movie models.Movie
				err := json.Unmarshal(reqBody, &movie)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				models.MapOfMovies[i][id] = movie
				err = json.NewEncoder(w).Encode(models.MapOfMovies)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
		if !found {
			log.Println("Movie not found")
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}
		_, err := w.Write([]byte("movie updated"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateMoviesInBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		log.Println("updating movies in batch with body: " + string(reqBody))
		var movies []models.Movie
		err := json.Unmarshal(reqBody, &movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var foundId []string
		for _, movie := range movies {
			for i, movieMap := range models.MapOfMovies {
				if _, ok := movieMap[movie.ID.String()]; ok {
					foundId = append(foundId, movie.ID.String())
					models.MapOfMovies[i][movie.ID.String()] = movie
				}
			}
		}
		if len(foundId) == 0 {
			log.Println("No movies found")
			http.Error(w, "No movies found", http.StatusNotFound)
			return
		}
		_, err = w.Write([]byte("movies updated"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		log.Println("deleting movie with id: " + id)
		found := false
		for i, movieMap := range models.MapOfMovies {
			if _, ok := movieMap[id]; ok {
				found = true
				delete(models.MapOfMovies[i], id)
			}
		}
		if !found {
			log.Println("Movie not found")
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}
		_, err := w.Write([]byte("Movie deleted"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteMoviesInBatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)
		log.Println("deleting movies in batch with body: " + string(reqBody))
		var movies []models.Movie
		err := json.Unmarshal(reqBody, &movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var foundId []string
		for _, movie := range movies {
			for i, movieMap := range models.MapOfMovies {
				for key, value := range movieMap {
					if value.Title == movie.Title {
						foundId = append(foundId, key)
						delete(models.MapOfMovies[i], key)
					}
				}
			}
		}
		if len(foundId) == 0 {
			log.Println("No movies found")
			http.Error(w, "No movies found", http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(models.MapOfMovies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func SearchMovies() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("t")
		year := r.URL.Query().Get("y")
		dir := r.URL.Query().Get("d")
		imdb := r.URL.Query().Get("i")
		genre := r.URL.Query().Get("g")
		log.Println("searching movies with query: " + title)
		var movies []models.Movie
		moviesSet := make(map[string]models.Movie)
		for _, movieMap := range models.MapOfMovies {
			for _, movie := range movieMap {
				if title != "" && !strings.Contains(movie.Title, title) {
					continue
				}
				if year != "" && strconv.Itoa(movie.Year) != year {
					continue
				}
				if dir != "" && !strings.Contains(movie.Director, dir) {
					continue
				}
				if imdb != "" && strconv.FormatFloat(movie.ImdbRating, 'E', -1, 64) != imdb {
					continue
				}
				if genre != "" && movie.Genre != genre {
					continue
				}
				moviesSet[movie.ID.String()] = movie
			}
		}
		for _, movie := range moviesSet {
			movies = append(movies, movie)
		}
		if len(movies) == 0 {
			log.Println("No movies found")
			http.Error(w, "No movies found", http.StatusNotFound)
			return
		}
		err := json.NewEncoder(w).Encode(movies)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
