package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	var err error
	if len(cmd.Args) == 0 {
		limit = 2
	} else if len(cmd.Args) == 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err) // Обработка ошибки
		}
	} else {
		return fmt.Errorf("usage: %s [limit]", cmd.Name)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf(" *   Post ID: %s\n", post.ID)
		fmt.Printf(" *   Published at: %s\n", post.PublishedAt.Time.Format("2006-01-02 15:04:05"))
		fmt.Printf(" *   URL: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Println("=====================================")
	}
	return nil
}
