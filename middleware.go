package main

import (
	"fmt"
	"net/http"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
)

type authedHandler func(http.ResponseWriter,*http.Request,database.User)

func (apiCfg *apiConfig) middleware(handler authedHandler) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		apiKey, err := GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error:%v",err))
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(),apiKey)

		if err != nil {
			respondWithError(w,403,fmt.Sprintf("Auth error:%v",err))
			return
		}

		handler(w, r , user)
	}
}