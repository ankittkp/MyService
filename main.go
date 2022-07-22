package main

import (
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"MyService/auth"
	"MyService/handlers"
)

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
			return "", fileName
		},
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Println("Starting My Service")
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/token", auth.CreateToken).Methods("POST")
	router.HandleFunc("/api/v1/movies", auth.Middleware(handlers.GetAllMovies)).Methods("GET")
	router.HandleFunc("/api/v1/movie/{id}", auth.Middleware(handlers.GetMovie)).Methods("GET")
	router.HandleFunc("/api/v1/movie", auth.Middleware(handlers.AddMovie())).Methods("POST")
	router.HandleFunc("/api/v1/movies/batch", auth.Middleware(handlers.AddMoviesInBatch())).Methods("POST")
	router.HandleFunc("/api/v1/movie/{id}", auth.Middleware(handlers.UpdateMovie())).Methods("PUT")
	router.HandleFunc("/api/v1/movies/batch", auth.Middleware(handlers.UpdateMoviesInBatch())).Methods("PUT")
	router.HandleFunc("/api/v1/movie/{id}", auth.Middleware(handlers.DeleteMovie())).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/batch", auth.Middleware(handlers.DeleteMoviesInBatch())).Methods("DELETE")
	router.HandleFunc("/api/v1/movies/search", auth.Middleware(handlers.SearchMovies())).Queries("t", "{t}").Queries("y", "{y}").Queries("d", "{d}").Queries("i", "{i}").Queries("g", "{g}").Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
