package http

import (
	"context"
	"errors"
	"strings"
)

// mutation.go is the place where we want to implement mutation functionality

func (tR *twitterHTTPRepo) PostTweet(ctx context.Context, tweet string) (err error) {
	if strings.TrimSpace(tweet) == "" {
		err = errors.New("empty tweet")
		return
	}

	if tR.tClient == nil {
		err = errors.New("no twitter instance")
		return
	}

	_, _, err = tR.tClient.Statuses.Update(tweet, nil)
	return
}
