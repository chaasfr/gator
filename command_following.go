package main

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {

	user, err := s.dbQueries.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching current user: %w", err)
	}

	feeds, err := s.dbQueries.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching follows for user %s: %w", user.Name, err)
	}

	for _, feed := range feeds {
		fmt.Printf("- %s\n", feed.Feedname)
	}

	return nil
}