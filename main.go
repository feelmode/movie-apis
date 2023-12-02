package main

import (
	"encoding/json"
	"log"
	movieHttpHandler "main/internal/movie/http/handler"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const baseMoviePath = "/movies"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	h := &movieHttpHandler.Handler{}
	h.Db = getDb()

	r.HandleFunc(baseMoviePath, h.GetHandler).Methods("GET")
	r.HandleFunc(baseMoviePath, h.PostHandler).Methods("POST")
	r.HandleFunc(baseMoviePath+"/{id}", h.DeleteByIDHandler).Methods("DELETE")
	r.HandleFunc(baseMoviePath+"/{id}", h.GetByIDHandler).Methods("GET")
	r.HandleFunc(baseMoviePath+"/{id}", h.PatchByIDHandler).Methods("PATCH")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getDb() *gorm.DB {
	dsn := "host=localhost user=postgres password='' dbname=movies port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
