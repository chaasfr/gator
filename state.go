package main

import (
	"database/sql"

	"github.com/chaasfr/gator/internal/config"
	"github.com/chaasfr/gator/internal/database"
)

type State struct {
	conf      *config.Config
	dbQueries *database.Queries
	
}

func InitState() (*State, error){
	conf, err := config.Read()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", conf.DBURL)
	if err != nil {
		return nil, err
	}

	dbQueries := database.New(db)

	var state State
	state.conf = conf
	state.dbQueries = dbQueries
	return &state, nil
}