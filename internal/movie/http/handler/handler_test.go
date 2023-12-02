package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/internal/movie"
	gormDB "main/pkg/database/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

const baseMoviePath = "/movie"
const fakeID = "99"
const pathWithID = baseMoviePath + "/" + fakeID

var h *Handler

func init() {
	h = &Handler{}
	h.Db = gormDB.GetDB()
	h.Db.Exec("DELETE FROM movies")
	h.Db.Exec("ALTER SEQUENCE movies_id_seq RESTART WITH 1")
}

func TestCreateResp(t *testing.T) {
	result := createResp(movie.Movie{
		ID:          1,
		Title:       "Title 1",
		Description: "Description 1",
		Rating:      10,
		Image:       "image1.jpg",
	})

	expected := movie.Response{
		ID:          1,
		Title:       "Title 1",
		Description: "Description 1",
		Rating:      10,
		Image:       "image1.jpg",
	}

	if expected != result {
		t.Errorf("result was incorrect, got: %v, want: %v.", result, expected)
	}
}

func TestPostHandler(t *testing.T) {
	var jsonStr = []byte(`{"title": "Title 1", "description": "Desc 1"}`)
	req, _ := http.NewRequest("POST", baseMoviePath, bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostHandlerBadRequestNoBody(t *testing.T) {
	var jsonStr = []byte(``)
	req, _ := http.NewRequest("POST", baseMoviePath, bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestPatchByIDHandlerBadRequestNoBody(t *testing.T) {
	var jsonStr = []byte(``)
	req, _ := http.NewRequest("PATCH", pathWithID, bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": fakeID,
	})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PatchByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestPostHandlerBadRequest(t *testing.T) {
	var jsonStr = []byte(`{"rating": 7, "image": "image1.jpg"}`)
	req, _ := http.NewRequest("POST", baseMoviePath, bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	if !strings.Contains(rr.Body.String(), "ERR__BAD_REQUEST_FIELDS") {
		t.Errorf("body not containing ERR__BAD_REQUEST_FIELDS")
	}
}

func TestGetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", baseMoviePath, nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetByIDHandler(t *testing.T) {
	id, _, rr := getNewlyCreatedID()
	req, _ := http.NewRequest("GET", baseMoviePath+"/"+id, nil)

	// See https://stackoverflow.com/questions/34435185/unit-testing-for-functions-that-use-gorilla-mux-url-parameters
	req = mux.SetURLVars(req, map[string]string{
		"id": id,
	})

	handler := http.HandlerFunc(h.GetByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetByIDHandlerNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", pathWithID, nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": fakeID,
	})
	handler := http.HandlerFunc(h.GetByIDHandler)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestPatchHandler(t *testing.T) {
	id, _, rr := getNewlyCreatedID()
	jsonStr := []byte(`{"title": "Title 1a", "description": "Desc 1a"}`)
	req, _ := http.NewRequest("PATCH", baseMoviePath+"/"+id, bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": id,
	})
	handler := http.HandlerFunc(h.PatchByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func getNewlyCreatedID() (string, Handler, *httptest.ResponseRecorder) {
	var jsonStr = []byte(`{"title": "Title 1", "description": "Desc 1"}`)
	req, _ := http.NewRequest("POST", baseMoviePath, bytes.NewBuffer(jsonStr))
	h := *h
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PostHandler)
	handler.ServeHTTP(rr, req)
	var resp map[string]map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)
	id := fmt.Sprintf("%v", resp["data"]["id"])

	return id, h, rr
}

func TestDeleteByIDHandler(t *testing.T) {
	id, _, _ := getNewlyCreatedID()
	req, _ := http.NewRequest("GET", baseMoviePath+"/"+id, nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": id,
	})

	// Must regenerate these stuff to get the actual data
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.DeleteByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}

func TestDeleteByIDHandlerNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", pathWithID, nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": fakeID,
	})

	// Must regenerate these stuff to get the actual data
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(h.DeleteByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestPatchHandlerBadRequest(t *testing.T) {
	id, _, _ := getNewlyCreatedID()
	jsonStr := []byte(`{"rating": 9}`)
	req, _ := http.NewRequest("PATCH", baseMoviePath+"/"+id, bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": id,
	})
	handler := http.HandlerFunc(h.PatchByIDHandler)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestPatchHandlerNotFound(t *testing.T) {
	jsonStr := []byte(`{"title": "Title 1a", "description": "Desc 1a"}`)
	req, _ := http.NewRequest("PATCH", pathWithID, bytes.NewBuffer(jsonStr))
	req = mux.SetURLVars(req, map[string]string{
		"id": fakeID,
	})
	handler := http.HandlerFunc(h.PatchByIDHandler)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
