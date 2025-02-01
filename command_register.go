package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chaasfr/gator/internal/database"
	"github.com/google/uuid"
)



func handlerRegister(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error - please provide a name to register")
	}
	username := cmd.args[0]
	qp := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username,
	}
	user, err := s.dbQueries.CreateUser(context.Background(),qp)
	if err != nil {
		return fmt.Errorf("error creating user %s. %w", username, err)
	}

	s.conf.SetUser(user.Name)

	fmt.Printf("user created: %s\n", user)

	return nil
}