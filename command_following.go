package main

import (
	"context"
	"fmt"

	"github.com/chaasfr/gator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {

	feeds, err := s.dbQueries.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching follows for user %s: %w", user.Name, err)
	}

	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.Feedname)
	}

	return nil
}