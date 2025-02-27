package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctxt := context.Background()
	err := s.db.DeleteUsers(ctxt)
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
