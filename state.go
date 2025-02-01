package main

import (
	"github.com/chaasfr/gator/internal/config"
)

type State struct {
	conf *config.Config
}

func InitState() (*State, error){
	conf, err := config.Read()
	if err != nil {
		return nil, err
	}
	var state State
	state.conf = conf
	return &state, nil
}