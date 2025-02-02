package main

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	
	feeds, err := s.dbQueries.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%s: %s - %s\n", feed.Name, feed.Url, feed.Username.String)
	}

	return nil
}