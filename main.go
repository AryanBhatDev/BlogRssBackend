package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct{
	DB *database.Queries
}


func main(){

	feed, err := urlToFeed("https://wagslane.dev/index.xml")

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(feed)

	err = godotenv.Load(".env")

	if err != nil{
		log.Fatal("Cannot find .env file")
	}

	port := os.Getenv("PORT")

	if port == ""{
		log.Fatal("PORT not found in .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")

	if dbUrl == ""{
		log.Fatal("DATABASE_URL not found in .env file")
	}

	conn, err := sql.Open("postgres",dbUrl)

	if err != nil {
		log.Fatal("Error while connecting to database",err)
	}
	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	go scrapperForever(queries,10,time.Minute)

	router := chi.NewRouter()
	
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods:  []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge:			300,
	}))


	v1Router := chi.NewRouter()

	v1Router.Get("/ready",handlerReadiness)
	v1Router.Get("/err",handleErr)
	v1Router.Post("/user",apiCfg.handlerCreateUser)
	v1Router.Get("/user",apiCfg.middleware(apiCfg.handlerGetUser))
	v1Router.Post("/feed",apiCfg.middleware(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed/all",apiCfg.handlerGetFeeds)
	v1Router.Post("/user/feed/follow",apiCfg.middleware(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/user/feed/follows/all",apiCfg.middleware(apiCfg.handlerGetFeedFollows))
	v1Router.Post("/user/feed/follow/{feedFollowId}",apiCfg.middleware(apiCfg.handlerDeleteFeedFollow))
	v1Router.Post("/user/posts",apiCfg.middleware(apiCfg.handlerGetPostsForUser))
	router.Mount("/api/v1",v1Router)

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