package http

import (
	"github.com/dghubble/go-twitter/twitter"
	tR "github.com/robertotambunan/tweather/repositories/twitter"
)

const (
	// TwitterV1Endpoint : twitter v1 endpoint json
	TwitterV1Endpoint = "https://api.twitter.com/1.1/statuses/update.json"
)

type twitterHTTPRepo struct {
	tClient *twitter.Client
}

// NewTwitterHTTP : Initialize Twitter Repo
func NewTwitterHTTP(twitterClient *twitter.Client) tR.Repository {
	return &twitterHTTPRepo{
		tClient: twitterClient,
	}
}
