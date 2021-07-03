package news

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	nR "github.com/robertotambunan/tweather/repositories/news"
	tR "github.com/robertotambunan/tweather/repositories/twitter"
)

func (nu *newsUC) PostTopNewsIndonesia(ctx context.Context) (err error) {
	// id -> news for indonesia
	topNews := nu.newsRepo.GetTopNewsBasedOnNation(ctx, "id", "", 5)
	if len(topNews) <= 0 {
		err = errors.New("No top news available")
		return
	}

	lastID := int64(0)
	var tweetingErr error
	idx := 1
	for _, v := range topNews {
		tweetPost := constructNewsToTweet(v, idx)

		// having a watermark in first tweet
		if lastID == 0 {
			tweetPost += " - by Roberto BOT"
		}

		// note: first tweet will have lastID 0 - it's okay, we cater it in repository validation
		twtID, errTwt := nu.twitterRepo.PostTweet(ctx, tR.TweetParam{
			Tweet:   tweetPost,
			ReplyTo: lastID,
		})

		if errTwt != nil {
			// collecting error in one variable
			prevErr := ""
			if tweetingErr != nil {
				prevErr = tweetingErr.Error()
			}
			tweetingErr = errors.New(prevErr + " " + errTwt.Error())
		} else {
			idx++
			lastID = twtID
		}
		// having a break to next tweet
		time.Sleep(500 * time.Millisecond)
	}
	err = tweetingErr
	return
}

func (nu *newsUC) PostTopNewsSport(ctx context.Context) (err error) {
	// id -> news for indonesia
	topNews := nu.newsRepo.GetTopNewsBasedOnNation(ctx, "gb", "sports", 5)
	if len(topNews) <= 0 {
		err = errors.New("No top news available")
		return
	}

	lastID := int64(0)
	var tweetingErr error
	idx := 1
	for _, v := range topNews {
		tweetPost := constructNewsToTweet(v, idx)

		// having a watermark in first tweet
		if lastID == 0 {
			tweetPost += " - by Roberto BOT"
		}

		// note: first tweet will have lastID 0 - it's okay, we cater it in repository validation
		twtID, errTwt := nu.twitterRepo.PostTweet(ctx, tR.TweetParam{
			Tweet:   tweetPost,
			ReplyTo: lastID,
		})

		if errTwt != nil {
			// collecting error in one variable
			prevErr := ""
			if tweetingErr != nil {
				prevErr = tweetingErr.Error()
			}
			tweetingErr = errors.New(prevErr + " " + errTwt.Error())
		} else {
			idx++
			lastID = twtID
		}
		// having a break to next tweet
		time.Sleep(500 * time.Millisecond)
	}
	err = tweetingErr
	return
}

func constructNewsToTweet(item nR.News, index int) (formatted string) {
	formatted = strconv.Itoa(index) + ") " + strings.Title(item.Title) + " - " + item.URL
	return
}
