package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/google/uuid"
)



func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)


	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while decoding: %v",err))
		return
	}
	user, err := apiCfg.DB.CreateUser(
		r.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      params.Name,
			Email:     params.Email,
			Password:  params.Password,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating user: %v",err))
		return
	}

	respondWithJson(w,200,databaseUserToUser(user))
}

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request){
	
	apiKey, err := GetApiKey(r.Header)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error:%v",err))
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(),apiKey)

	if err != nil {
		respondWithError(w,403,fmt.Sprintf("Auth error:%v",err))
		return
	}

	respondWithJson(w,200,databaseUserToUser(user))
}