package main

import "github.com/AryanBhatDev/blogrssbackend/internal/database"


type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ApiKey string `json:apiKey`
}


type GetUser struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
type Feed struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type FeedFollow struct {
	UserID string `json:"user_id"`
	FeedID string 	`json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow{
	return FeedFollow{
		UserID : dbFeedFollow.UserID.String(),
		FeedID : dbFeedFollow.FeedID.String(),
	}
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
func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow{

	feedFollows := []FeedFollow{}

	for _,feedFollow := range dbFeedFollows{
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}


func databaseUserToGetUser(dbUser database.User) GetUser{
	return GetUser{
		Name:dbUser.Name,
		Email: dbUser.Email,
	}
}

