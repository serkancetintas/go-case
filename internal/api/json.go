package api

import (
	"encoding/json"
	"net/http"
)

func JsonErr(w http.ResponseWriter, status int, err error) {
	Json(w, status, map[string]string{"error": err.Error()})
}

func Json(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
