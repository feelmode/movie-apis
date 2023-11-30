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

func createResp(req movie.Movie) movie.Response {
	return movie.Response{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
		CreatedAt:   req.CreatedAt,
		UpdatedAt:   req.UpdatedAt,
	}
}

func getDb() *gorm.DB {
	dsn := "host=localhost user=qosdil password='' dbname=movies port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error: " + err.Error())
		// resp.Write(w, http.StatusInternalServerError, nil, nil)
	}

	return db
}

func DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	var movie movie.Movie
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db := getDb()
	db.First(&movie, id)

	// Not found
	if movie.ID == 0 {
		resp.Write(w, http.StatusNotFound, nil, nil)
		return
	}

	getDb().Delete(&movie, id)
	resp.Write(w, http.StatusNoContent, nil, nil)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	response := []movie.Response{}
	db := getDb()
	rows, _ := db.Find(&[]movie.Movie{}).Rows()
	defer rows.Close()
	for rows.Next() {
		var movie movie.Movie
		db.ScanRows(rows, &movie)
		response = append(response, createResp(movie))
	}

	resp.Write(w, http.StatusOK, nil, response)
}

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	var movie movie.Movie
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	getDb().First(&movie, id)
	if movie.ID == 0 {
		resp.Write(w, http.StatusNotFound, nil, nil)
		return
	}

	resp.Write(w, http.StatusOK, nil, createResp(movie))
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	var req movie.Movie
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		resp.Write(w, http.StatusBadRequest, &resp.Error{Code: resp.ERR__BAD_REQUEST_FIELDS, Message: err.Error()}, nil)
		return
	}

	var movie movie.Movie
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	db := getDb()
	db.First(&movie, id)

	// Not found
	if movie.ID == 0 {
		resp.Write(w, http.StatusNotFound, nil, nil)
		return
	}

	// Save
	movie.Title = req.Title
	movie.Description = req.Description
	movie.Rating = req.Rating
	movie.Image = req.Image
	db.Save(&movie)

	resp.Write(w, http.StatusOK, nil, createResp(movie))
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

	getDb().Create(&reqResp)
	resp.Write(w, http.StatusOK, nil, createResp(reqResp))
}
