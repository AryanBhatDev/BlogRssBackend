package main

import "github.com/AryanBhatDev/blogrssbackend/internal/database"


type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ApiKey string `json:apiKey`
}

type Feed struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type GetUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func databaseUserToUser(dbUser database.User) User{
	return User{
		Name:dbUser.Name,
		Email: dbUser.Email,
		ApiKey: dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed{
	return Feed{
		Name:dbFeed.Name,
		Url:dbFeed.Url,
	}
}

func databaseFeedsToFeeds(dbFeed []database.Feed) []Feed{

	feeds := []Feed{}

	for _,feed := range dbFeed{
		feeds = append(feeds, databaseFeedToFeed(feed))
	}
	return feeds
}

func databaseUserToGetUser(dbUser database.User) GetUser{
	return GetUser{
		Name:dbUser.Name,
		Email: dbUser.Email,
	}
}

