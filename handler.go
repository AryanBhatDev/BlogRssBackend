package main

import "net/http"


func handlerReadiness(w http.ResponseWriter, r *http.Request){
	respondWithJson(w,200,struct{}{})
}

func handleErr(w http.ResponseWriter, r *http.Request){
	respondWithError(w,400,"Something went wrong")
}