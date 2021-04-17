package twitter

import "context"

// Repository contract for twitter repository
type Repository interface {
	PostTweet(ctx context.Context, tweet string) (err error)
}
