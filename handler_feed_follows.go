package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/go-chi/chi"
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

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollows , err := apiCfg.DB.GetFeedFollows(r.Context(),user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follows:%v",err))
		return
	}
	respondWithJson(w,200,databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){

	feedFollowIdStr := chi.URLParam(r,"feedFollowId")

	FeedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id:%v",err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: FeedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow:%v",err))
		return 
	}

	respondWithJson(w,200,struct{}{})
}