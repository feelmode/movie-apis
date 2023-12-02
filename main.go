package main

import (
	"encoding/json"
	"log"
	movieHttpHandler "main/internal/movie/http/handler"
	gormDB "main/pkg/database/gorm"
	"net/http"

	"github.com/gorilla/mux"
)

const baseMoviePath = "/movies"
const moviePathWithID = baseMoviePath + "/{id}"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	h := &movieHttpHandler.Handler{}
	h.Db = gormDB.GetDB()

	r.HandleFunc(baseMoviePath, h.GetHandler).Methods("GET")
	r.HandleFunc(baseMoviePath, h.PostHandler).Methods("POST")
	r.HandleFunc(moviePathWithID, h.DeleteByIDHandler).Methods("DELETE")
	r.HandleFunc(moviePathWithID, h.GetByIDHandler).Methods("GET")
	r.HandleFunc(moviePathWithID, h.PatchByIDHandler).Methods("PATCH")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
