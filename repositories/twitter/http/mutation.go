package http

import (
	"context"
	"errors"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	tR "github.com/robertotambunan/tweather/repositories/twitter"
)

// mutation.go is the place where we want to implement mutation functionality

func (tR *twitterHTTPRepo) PostTweet(ctx context.Context, twtParam tR.TweetParam) (tweetID int64, err error) {
	if strings.TrimSpace(twtParam.Tweet) == "" {
		err = errors.New("empty tweet")
		return
	}

	if tR.tClient == nil {
		err = errors.New("no twitter instance")
		return
	}

	tweetParam := &twitter.StatusUpdateParams{}
	if twtParam.ReplyTo > 0 {
		tweetParam.InReplyToStatusID = twtParam.ReplyTo
	} else {
		tweetParam = nil
	}

	t := &twitter.Tweet{}
	t, _, err = tR.tClient.Statuses.Update(twtParam.Tweet, tweetParam)

	if t != nil {
		tweetID = t.ID
	}

	return
}
