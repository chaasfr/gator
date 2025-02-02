package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chaasfr/gator/internal/database"
)


func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		i, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("error parsing %s, please provide an int. %w", cmd.args[0], err)
		}
		limit = i
	}
	
	qpGetPosts := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.dbQueries.GetPostsForUser(context.Background(), qpGetPosts)

	if err != nil {
		return fmt.Errorf("error geting %v posts for %s: %w", limit, user.Name, err)
	}

	for _, post := range posts {
		fmt.Println("================================")
		fmt.Println("Title: " + post.Title.String)
		fmt.Println("Link: " + post.Url)
		fmt.Printf("published at: %v \n", post.PublishedAt)
		fmt.Println("description: " + post.Description.String)
	}

	return nil
}