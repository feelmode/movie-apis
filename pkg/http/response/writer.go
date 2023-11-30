package writer

import (
	"encoding/json"
	"net/http"
)

func Write(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if resp != nil {
		json.NewEncoder(w).Encode(resp)
	}
}
