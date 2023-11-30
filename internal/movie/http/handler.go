package handler

import (
	"encoding/json"
	"main/internal/movie"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	movies := []movie.Resp{}
	movies = append(movies, movie.Resp{
		ID:          1,
		Title:       "Title 1",
		Description: "Desc 1",
		Rating:      7,
		Image:       "",
	})
	movies = append(movies, movie.Resp{
		ID:          2,
		Title:       "Title 2",
		Description: "Desc 3",
		Rating:      9,
		Image:       "",
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
