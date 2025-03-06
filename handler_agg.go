package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/koldunNomad/gator/internal/database"

	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("incorrect duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s", cmd.Args[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

// Печатает feeds пользователя
func scrapeFeeds(s *state) error {
	// Получает следующий feed
	ctxt := context.Background()

	feedData, err := s.db.GetNextFeedToFetch(ctxt)
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %w", err)
	}

	// Отмечает его как полученый
	err = s.db.MarkFeedFetched(ctxt, feedData.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	// Получаем feed по url
	fetchedFeed, err := fetchFeed(ctxt, feedData.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch a feed: %w", err)
	}

	for _, item := range fetchedFeed.Channel.Item {
		fmt.Println(" * ", item.Title)

		pulishedAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("Error parsing published date:", err)
			continue
		}

		description := sql.NullString{
			String: item.Description,
			Valid:  item.Description != "",
		}
		publishedAtNull := sql.NullTime{
			Time:  pulishedAt,
			Valid: true,
		}

		err = s.db.CreatePost(ctxt, database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAtNull,
			FeedID:      feedData.ID,
		})

		if err != nil {
			log.Println("Error inserting post:", err)
			continue
		}
	}

	return nil
}
