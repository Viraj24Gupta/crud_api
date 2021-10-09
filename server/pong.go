package server

import (
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(`{"success": "pong"}`)); err != nil {
	return
	}
}