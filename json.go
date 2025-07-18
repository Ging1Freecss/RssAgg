package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter,code int, msg string){

	if code >= 500 {
		log.Println("Responding with 5XX error : ",msg);
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w,code,errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	

	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}