package handler

import (
	"encoding/json"
	"main/internal/movie"
	resp "main/pkg/http/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	Db *gorm.DB
}

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

func (h Handler) DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	var movie movie.Movie
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	h.Db.First(&movie, id)

	// Not found
	if movie.ID == 0 {
		resp.Write(w, http.StatusNotFound, nil, nil)
		return
	}

	h.Db.Delete(&movie, id)
	resp.Write(w, http.StatusNoContent, nil, nil)
}

func (h Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	response := []movie.Response{}
	rows, _ := h.Db.Find(&[]movie.Movie{}).Rows()
	defer rows.Close()
	for rows.Next() {
		var movie movie.Movie
		h.Db.ScanRows(rows, &movie)
		response = append(response, createResp(movie))
	}

	resp.Write(w, http.StatusOK, nil, response)
}

func (h Handler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	var movie movie.Movie
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	h.Db.First(&movie, id)
	if movie.ID == 0 {
		resp.Write(w, http.StatusNotFound, nil, nil)
		return
	}

	resp.Write(w, http.StatusOK, nil, createResp(movie))
}

func (h Handler) PatchHandler(w http.ResponseWriter, r *http.Request) {
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
	h.Db.First(&movie, id)

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
	h.Db.Save(&movie)

	resp.Write(w, http.StatusOK, nil, createResp(movie))
}

func (h Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
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

	h.Db.Create(&reqResp)
	resp.Write(w, http.StatusOK, nil, createResp(reqResp))
}
