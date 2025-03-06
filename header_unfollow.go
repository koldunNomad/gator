package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func headerunfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	err := s.db.DeleteFollowRecord(context.Background(), database.DeleteFollowRecordParams{
		Name: user.Name,
		Url:  cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("delete querie error: %w", err)
	}

	fmt.Println("Unfollow completed")
	return nil
}
