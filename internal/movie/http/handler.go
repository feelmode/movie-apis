package handler

import (
	"encoding/json"
	"main/internal/movie"
	resp "main/pkg/http/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	resp.Write(w, http.StatusNoContent, resp.NO_ERROR, nil)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	movies := []movie.ReqResp{}
	movies = append(movies, movie.ReqResp{
		ID:          1,
		Title:       "Title 1",
		Description: "Desc 1",
		Rating:      7,
		Image:       "",
	})
	movies = append(movies, movie.ReqResp{
		ID:          2,
		Title:       "Title 2",
		Description: "Desc 3",
		Rating:      9,
		Image:       "",
	})

	resp.Write(w, http.StatusOK, resp.NO_ERROR, movies)
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	movie := movie.ReqResp{
		ID:          uint8(id),
		Title:       "Title 1",
		Description: "Desc 1",
		Rating:      7,
		Image:       "",
	}

	resp.Write(w, http.StatusOK, resp.NO_ERROR, movie)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	var reqResp movie.ReqResp
	err := json.NewDecoder(r.Body).Decode(&reqResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Write(w, http.StatusOK, resp.NO_ERROR, reqResp)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqResp movie.ReqResp
	err := json.NewDecoder(r.Body).Decode(&reqResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp.Write(w, http.StatusOK, resp.NO_ERROR, reqResp)
}
