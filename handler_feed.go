package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/google/uuid"
)



func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request,user database.User){
	type parameters struct{
		Url string `json:"url"`
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)


	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while decoding: %v",err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(
		r.Context(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      params.Name,
			Url:	params.Url,
			UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating user: %v",err))
		return
	}

	respondWithJson(w,200,databaseFeedToFeed(feed))
}