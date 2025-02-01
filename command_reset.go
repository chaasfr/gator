package main

import "context"

func HandlerReset(s *State, cmd Command) error {
	if err := s.dbQueries.DeleteUsers(context.Background()); err != nil {
		return err
	}

	return nil
}