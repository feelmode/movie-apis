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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	h := &movieHttpHandler.Handler{}
	h.Db = getDb()

	r.HandleFunc("/movies", h.GetHandler).Methods("GET")
	r.HandleFunc("/movies", h.PostHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", h.DeleteByIDHandler).Methods("DELETE")
	r.HandleFunc("/movies/{id}", h.GetByIDHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", h.PatchHandler).Methods("PATCH")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getDb() *gorm.DB {
	dsn := "host=localhost user=qosdil password='' dbname=movies port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
