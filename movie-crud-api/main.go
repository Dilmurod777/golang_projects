package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie
var directors []Director

func getMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)

	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)

			err := json.NewEncoder(w).Encode(movies)
			if err != nil {
				return
			}
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	directors = append(directors, Director{Firstname: "John", Lastname: "Doe"})
	directors = append(directors, Director{Firstname: "Ashley", Lastname: "Smith"})

	movies = append(movies, Movie{ID: "1", ISBN: "438226", Title: "Movie one", Director: &directors[0]})
	movies = append(movies, Movie{ID: "2", ISBN: "169874", Title: "Movie one", Director: &directors[1]})
	movies = append(movies, Movie{ID: "3", ISBN: "789926", Title: "Movie one", Director: &directors[1]})
	movies = append(movies, Movie{ID: "4", ISBN: "766631", Title: "Movie one", Director: &directors[0]})
	movies = append(movies, Movie{ID: "5", ISBN: "159846", Title: "Movie one", Director: &directors[1]})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}