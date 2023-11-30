package response

import (
	"encoding/json"
	"net/http"
)

const (
	ERR__BAD_REQUEST_FIELDS = "ERR__BAD_REQUEST_FIELDS"
)

type Response struct {
	Error *Error      `json:"error"` // "null" value possible
	Data  interface{} `json:"data"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Write(w http.ResponseWriter, status int, err *Error, data interface{}) (resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp.Data = data
	if err != nil {
		switch err.Code {
		default:
			resp.Error = &Error{
				Code:    err.Code,
				Message: err.Message,
			}
		}
	}

	json.NewEncoder(w).Encode(resp)
	return resp
}
