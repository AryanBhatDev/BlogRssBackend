package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/AryanBhatDev/blogrssbackend/internal/database"
	"github.com/google/uuid"
)

func scrapperForever(
	db *database.Queries, 
	concurrency int, 
	timeBetweenRequests time.Duration,
){
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C{
		feeds, err := db.GetNextFeedsToFetch(context.Background(),
			int32(concurrency),
		)
		if err != nil{
			log.Println("error fetching feeds:",err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds{
			wg.Add(1)

			go scrapeFeed(db,wg,feed)

		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries,wg *sync.WaitGroup,feed database.Feed){
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(),feed.ID)

	if err != nil{
		log.Println("err marking feed as fetched:",err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed:",err)
		return
	}

	for _,item := range rssFeed.Channel.Item{
		description := sql.NullString{}

		if item.Description != ""{
			description.String = item.Description
			description.Valid = true
		}

		parsedTime, err := parseAnyTime(item.PubDate)

		if err != nil {
			log.Println("Error parsing time:", err)
			continue
		} 

		_, err = db.CreatePosts(context.Background(),database.CreatePostsParams{
			ID: uuid.New(),
			CreatedAt:time.Now(),
			UpdatedAt:time.Now(),
			Title: item.Title,
			Description: description,
			PublishedAt: parsedTime,
			Url: item.Link,
			FeedID: feed.ID,

		})
		if err != nil{
			if strings.Contains(err.Error(),"duplicate key"){
				return
			}
			log.Println("failed to create post",err)
			return
		}
	}


	log.Printf("Feed %s collected, %v posts found",feed.Name,len(rssFeed.Channel.Item))
}


func parseAnyTime(dateStr string) (time.Time, error) {

    layouts := []string{
        time.RFC1123Z,
        time.RFC1123,
        time.RFC822Z,
        time.RFC822,
        time.RFC3339,
        time.RFC3339Nano,
        time.ANSIC,
        time.UnixDate,
        time.RubyDate,
        time.Kitchen,
        "2006-01-02 15:04:05",      
        "02 Jan 2006 15:04:05 MST",
        "Mon, 02 Jan 2006 15:04:05 MST",
    }

    var t time.Time
    var err error
    for _, layout := range layouts {
        t, err = time.Parse(layout, dateStr)
        if err == nil {
            return t, nil
        }
    }
    return time.Time{}, fmt.Errorf("unable to parse time: %v", dateStr)
}