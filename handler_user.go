package main

import (
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

func (apiCfg *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	
	respondWithJson(w,200,databaseUserToGetUser(user))
}

func (apiCfg *apiConfig)handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User){
	posts , err := apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID:user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Issue getting user posts:%v",err))
		return
	}

	respondWithJson(w,200,databasePostsToPosts(posts))
}