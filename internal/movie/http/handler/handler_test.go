package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/internal/movie"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	dsn := "host=localhost user=postgres password='' dbname=movies port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
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
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	h := Handler{}
	h.Db = getDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PostHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPostHandlerBadRequest(t *testing.T) {
	var jsonStr = []byte(`{"rating": 7, "image": "image1.jpg"}`)
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	h := Handler{}
	h.Db = getDb()

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
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := Handler{}
	h.Db = getDb()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.GetHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetByIDHandler(t *testing.T) {
	id, h, rr := getNewlyCreatedID()
	req, _ := http.NewRequest("GET", "/movies/"+id, nil)
	handler := http.HandlerFunc(h.GetByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestPatchHandler(t *testing.T) {
	id, h, rr := getNewlyCreatedID()
	jsonStr := []byte(`{"title": "Title 1a", "description": "Desc 1a"}`)
	req, _ := http.NewRequest("PATCH", "/movies/"+id, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(h.PatchByIDHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func getNewlyCreatedID() (string, Handler, *httptest.ResponseRecorder) {
	var jsonStr = []byte(`{"title": "Title 1", "description": "Desc 1"}`)
	req, _ := http.NewRequest("POST", "/movies", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	h := Handler{}
	h.Db = getDb()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.PostHandler)
	handler.ServeHTTP(rr, req)
	var resp map[string]map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &resp)
	id := fmt.Sprintf("%v", resp["data"]["id"])

	return id, h, rr
}
