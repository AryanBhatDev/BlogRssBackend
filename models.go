package main

import (
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/google/uuid"
)


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

type Post struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
	Title string `json:"title"`
	Description *string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url string `json:"url"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databasePostToPost( dbPost database.Post) Post{

	var description *string
	if dbPost.Description.Valid{
		description = &dbPost.Description.String
	}
	return Post{
		Title: dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url: dbPost.Url,
		FeedID: dbPost.FeedID,
	}
}

func databasePostsToPosts( dbPosts []database.Post) []Post{
	posts := []Post{}

	for _, dbPost := range dbPosts{
		posts = append(posts, databasePostToPost(dbPost))
	}

	return posts
}