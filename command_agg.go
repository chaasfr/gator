package main

import (
	"context"
	"fmt"
	"time"

	"github.com/chaasfr/gator/internal/database"
	"github.com/chaasfr/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("too few arguments. Usage: agg [time_between_reqs]")
	}

	time_between_reqs := cmd.args[0]
	deltaTime, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("cannot parseduration %s: %w", time_between_reqs, err)
	}

	fmt.Println("collecting feeds every "+ deltaTime.String())

	ticker := time.NewTicker(deltaTime)

	for ; ; <- ticker.C{
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}


func scrapeFeeds(s *State) error {
	feed, err := s.dbQueries.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error retrieving next feed to fech: %w", err)
	}

	fmt.Printf("Fetching %s at %s\n", feed.Name, feed.Url)

	rssFeed, err := rss.Fetchfeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	qpMarkFeed := database.MarkFeedFetchedParams{
		ID: feed.ID,
		UpdatedAt: time.Now(),
	}
	err = s.dbQueries.MarkFeedFetched(context.Background(), qpMarkFeed)
	if err != nil {
		return fmt.Errorf("error marking feed %s as fetched: %w", feed.Name, err)
	}

	for _, item := range rssFeed.Channel.Item{
		fmt.Printf("- %s\n", item.Title)
	}

	return nil
}