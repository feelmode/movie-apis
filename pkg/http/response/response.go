package response

import (
	"encoding/json"
	"net/http"
)

const (
	NO_ERROR    = "" // Need this because caller can't send nil
	ERR_FOO_BAR = "FOO_BAR"
)

type Response struct {
	Error *Error      `json:"error"` // "null" value possible
	Data  interface{} `json:"data"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Write(w http.ResponseWriter, status int, errCode string, data interface{}) (resp Response) {
	w.Header().Set("Content-Type", "application/json")
	resp.Data = data
	switch errCode {
	case ERR_FOO_BAR:
		resp.Error = &Error{
			Code:    "FOO_BAR",
			Message: "Some foobar err",
		}
	}

	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(status)
	return resp
}
