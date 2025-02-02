package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chaasfr/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("too few argument - usage: addfeed [name] [url]")
	}

	queryParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.args[0],
		Url: cmd.args[1],
		UserID: user.ID,
	}

	feed, err := s.dbQueries.CreateFeed(context.Background(), queryParams)
	if err != nil {
		return fmt.Errorf("error adding feed: %w", err)
	}

	fmt.Println(feed)

	return CreateFollow(s, user, feed)
}