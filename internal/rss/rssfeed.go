package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

const UserAgent = "gator"

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func (rf *RSSFeed) unescapeStrings(){
	rf.Channel.Title = html.UnescapeString(rf.Channel.Title)
	rf.Channel.Description = html.UnescapeString(rf.Channel.Description)

	for i, item := range rf.Channel.Item{
		rf.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rf.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}
}

func Fetchfeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var reader io.Reader
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, reader)

	if err != nil {
		return nil, fmt.Errorf("error creating req to %s", feedURL)
	}

	req.Header.Set("User-Agent", UserAgent)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error querying %v", req)
	}
	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading results from %s", feedURL)
	}

	var rssFeed *RSSFeed

	if err := xml.Unmarshal(resBytes, &rssFeed); err != nil {
		return nil, fmt.Errorf("error parsing xml - %w", err)
	}
	rssFeed.unescapeStrings()

	return rssFeed, nil
}