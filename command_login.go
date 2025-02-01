package main

import (
	"context"
	"fmt"
	"strings"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error login: a username is required")
	}
	username := strings.Join(cmd.args, " ")

	user, err := s.dbQueries.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error loging in %s - user not found. %w", username, err)
	}

	if err := s.conf.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("user set to %s\n", username)

	return nil
}
