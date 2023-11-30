package handler

import (
	"encoding/json"
	"fmt"
	"main/internal/movie"
	resp "main/pkg/http/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	resp.Write(w, http.StatusNoContent, nil, nil)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	movies := []movie.Movie{}
	movies = append(movies, movie.Movie{
		ID:          1,
		Title:       "Title 1",
		Description: "Desc 1",
		Rating:      7,
		Image:       "",
	})
	movies = append(movies, movie.Movie{
		ID:          2,
		Title:       "Title 2",
		Description: "Desc 3",
		Rating:      9,
		Image:       "",
	})

	resp.Write(w, http.StatusOK, nil, movies)
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	movie := movie.Movie{
		ID:          uint8(id),
		Title:       "Title 1",
		Description: "Desc 1",
		Rating:      7,
		Image:       "",
	}

	resp.Write(w, http.StatusOK, nil, movie)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	var reqResp movie.Movie
	err := json.NewDecoder(r.Body).Decode(&reqResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(reqResp)
	if err != nil {
		resp.Write(w, http.StatusBadRequest, &resp.Error{Code: resp.ERR__BAD_REQUEST_FIELDS, Message: err.Error()}, nil)
		return
	}

	resp.Write(w, http.StatusOK, nil, reqResp)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqResp movie.Movie
	err := json.NewDecoder(r.Body).Decode(&reqResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(reqResp)
	if err != nil {
		resp.Write(w, http.StatusBadRequest, &resp.Error{Code: resp.ERR__BAD_REQUEST_FIELDS, Message: err.Error()}, nil)
		return
	}

	dsn := "host=localhost user=qosdil password='' dbname=movies port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error: " + err.Error())
		resp.Write(w, http.StatusInternalServerError, nil, nil)
	}

	db.Create(&reqResp)
	resp.Write(w, http.StatusOK, nil, reqResp)
}
