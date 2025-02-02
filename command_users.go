package main

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.dbQueries.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("error retrieving users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.conf.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)	
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}