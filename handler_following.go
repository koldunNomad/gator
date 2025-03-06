package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

// Показать подписки
func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("couldn't get follows for user: %w", err)
	}

	for _, follow := range follows {
		printFollows(follow)
		fmt.Println("=====================================")
	}
	return nil
}

// Подписаться на <url>
func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	ctxt := context.Background()
	user, err := s.db.GetUser(ctxt, s.cfg.Current_user_name)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeed(ctxt, url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	feedFollowNote, err := s.db.CreateFeedFollow(ctxt, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("createFeedFollow:: %w", err)
	}

	fmt.Println("Followed successfully!")
	printFollowCreateNote(feedFollowNote)

	return nil
}

func printFollowCreateNote(follow database.CreateFeedFollowRow) {
	fmt.Printf("%s has successfully subscribed to the %s\n", follow.UserName, follow.FeedName)
}

func printFollows(follow database.GetFeedFollowsForUserRow) {
	fmt.Printf("* ID:            %s\n", follow.ID)
	fmt.Printf("* User:          %s\n", follow.UserName)
	fmt.Printf("* Feed:          %s\n", follow.FeedsName)
	fmt.Printf("* URL:           %s\n", follow.Url)
	fmt.Printf("* Created:       %v\n", follow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", follow.UpdatedAt)
}
