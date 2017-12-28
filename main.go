package main

import (
	"fmt"
	"github.com/jzelinskie/geddit"
	"github.com/reujab/wallpaper"
	"math/rand"
	"os"
	"time"
)

const (
	useragent = "gedditAgent v1"
	subreddit = "gwcommentsonearthporn"
)

var (
	username       = os.Getenv("REDDIT_USERNAME")
	password       = os.Getenv("REDDIT_PASSWORD")
	defaultOptions = geddit.ListingOptions{
		Time:  geddit.ThisMonth,
		Limit: 25,
	}
)

type Client struct {
	session *geddit.LoginSession
}

func main() {
	// Setup the Reddit API client
	client := new(Client)
	session, err := geddit.NewLoginSession(username, password, useragent)
	client.session = session
	if err != nil {
		panic(err)
	}

	// Get a random submission from the top of X time (defined by defaultOptions)
	submission := client.getRandomSubmission()
	err = wallpaper.SetFromURL(submission.URL)
	if err != nil {
		panic(err)
	}

	// Convenience message
	fmt.Printf("Set wallpaper to: '%s'\n", submission.URL)
}

func (c *Client) getRandomSubmission() *geddit.Submission {
	top, _ := c.session.SubredditSubmissions(subreddit, geddit.TopSubmissions, defaultOptions)
	length := len(top)
	rand.Seed(time.Now().UTC().UnixNano())

	return top[rand.Int()%length]
}
