package main

import (
	"fmt"
	"strings"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("error login: a username is required")
	}
	username := strings.Join(cmd.args, " ")
	if err := s.conf.SetUser(username); err != nil {
		return err
	}
	fmt.Printf("user set to %s\n", username)

	return nil
}
