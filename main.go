package main

import (
	"encoding/json"
	"log"
	movieHttpHandler "main/internal/movie/http/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	r.HandleFunc("/movies", movieHttpHandler.GetHandler).Methods("GET")
	r.HandleFunc("/movies", movieHttpHandler.PatchHandler).Methods("PATCH")
	r.HandleFunc("/movies", movieHttpHandler.PostHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", movieHttpHandler.DeleteByIDHandler).Methods("DELETE")
	r.HandleFunc("/movies/{id}", movieHttpHandler.GetByIDHandler).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
