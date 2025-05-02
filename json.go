package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithJson(w http.ResponseWriter, code int , payload interface{}){
	data , err := json.Marshal(payload)

	if err!=nil {
		log.Println("Failed to marshal response: ",payload)
		w.WriteHeader(500)
		return 
	}

	w.Header().Add("Content-type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}


func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Responding with err",msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondWithJson(w, code , errResponse{
		Error: msg,
	})
}