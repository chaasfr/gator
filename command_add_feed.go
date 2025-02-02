package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chaasfr/gator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("too few argument - usage: addfeed [name] [url]")
	}
	user, err := s.dbQueries.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
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