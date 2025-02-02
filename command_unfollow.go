package main

import (
	"context"
	"fmt"

	"github.com/chaasfr/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("too few argument. Usage: unfollow [url]")
	}
	url := cmd.args[0]
	queryParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url: url,
	}
	ff, err := s.dbQueries.DeleteFeedFollow(context.Background(), queryParams)
	if err != nil {
		return fmt.Errorf("error unfollowing %s: %w", url, err)
	}

	if ff == 0 {
		fmt.Printf("%s does not follow %s\n", user.Name, url)
	} else {
		fmt.Printf("%s unfollowed %s\n", user.Name, url)
	}

	return nil
}