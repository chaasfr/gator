package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chaasfr/gator/internal/database"
	"github.com/google/uuid"
)

func CreateFollow(s *State, user database.User, feed database.Feed) error {
	queryParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}
	ff, err := s.dbQueries.CreateFeedFollow(context.Background(), queryParams)
	if err != nil {
		return fmt.Errorf("error creating follow: %w", err)
	}

	fmt.Printf("%s is now following %s", ff.Username, ff.Feedname)
	return nil
}

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("too few argument. Usage: follow [url]")
	}
	feedURL := cmd.args[0]

	feed, err := s.dbQueries.GetFeedIdFromUrl(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error feed %s not found: %w", feedURL, err)
	}

	return CreateFollow(s, user, feed)
}
