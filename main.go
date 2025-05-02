package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)


func main(){

	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal("Cannot find .env file")
	}

	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("PORT not found in .env file")
	}

	router := chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods:  []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:			300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:  ":" + port,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Hello",port)
}