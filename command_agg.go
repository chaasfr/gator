package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"strings"
	"github.com/chaasfr/gator/internal/database"
	"github.com/chaasfr/gator/internal/rss"
	"github.com/google/uuid"
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
		err := savePost(s, item, feed.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func savePost(s *State, ri rss.RSSItem, feedID uuid.UUID) error {
	publishedAt, err := time.Parse(time.RFC1123Z, ri.PubDate)
	
	if err != nil {
		return fmt.Errorf("error parsing publishedAt %s", ri.PubDate)
	}

	qpCreatePost := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		Title:       sql.NullString{String: ri.Title, Valid: true},
		Url:         ri.Link,
		Description: sql.NullString{String: ri.Description, Valid: true},
		PublishedAt: publishedAt,
		FeedID:      feedID,
	}

	post, err := s.dbQueries.CreatePost(context.Background(), qpCreatePost)
	if err != nil && strings.Contains(err.Error(), "duplicate key") {
		fmt.Printf("duplicate keys - %s\n", ri.Title)
		return nil
	}
	if err != nil {
		return fmt.Errorf("error saving post %s: %w", ri.Title, err)
	}

	fmt.Printf("saved post %s", post.Title.String)
	return nil
}