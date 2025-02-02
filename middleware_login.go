package main

import (
	"fmt"
	"context"
	"github.com/chaasfr/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	result := func(s *State, cmd Command) error {
		user, err := s.dbQueries.GetUser(context.Background(),  s.conf.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting current user: %w", err)
		}
		return handler(s, cmd, user)
	}
	
	return result
}