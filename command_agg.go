package main

import (
	"context"
	"fmt"

	"github.com/chaasfr/gator/internal/rss"
)


const feedUrl = "https://www.wagslane.dev/index.xml"

func HandlerAgg(s *State, cmd Command) error {

	rssFeed, err := rss.Fetchfeed(context.Background(), feedUrl)

	if err != nil {
		return err
	}
	fmt.Println(rssFeed)

	return nil
}