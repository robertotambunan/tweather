package twitter

import "context"

//TweetParam : param for posting a tweet
type TweetParam struct {
	Tweet   string
	ReplyTo int64
}

// Repository contract for twitter repository
type Repository interface {
	// PostTweet : functionality to post a tweet, can be as a reply using replyTo
	PostTweet(ctx context.Context, twtParam TweetParam) (tweetID int64, err error)
}
