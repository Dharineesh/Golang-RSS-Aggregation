package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithErr(w http.ResponseWriter, code int, errMsg string) {
	if code > 499 {
		log.Println("Responding with 5xx error: %v", errMsg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondwithJSON(w, code, errResponse{
		Error: errMsg,
	})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
