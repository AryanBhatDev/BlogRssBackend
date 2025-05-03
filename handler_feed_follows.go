package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/google/uuid"
)



func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request,user database.User){
	type parameters struct{
		FeedId uuid.UUID `json:"feedId"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)


	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while decoding: %v",err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(
		r.Context(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID : params.FeedId,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating feed follow: %v",err))
		return
	}

	respondWithJson(w,200,databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request){
	feed , err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds:%v",err))
		return
	}
	respondWithJson(w,200,databaseFeedsToFeeds(feed))
}