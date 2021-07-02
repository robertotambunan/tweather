package news

import (
	nR "github.com/robertotambunan/tweather/repositories/news"
	tR "github.com/robertotambunan/tweather/repositories/twitter"
)

type newsUC struct {
	twitterRepo tR.Repository
	newsRepo    nR.Repository
}

// NewNewsUC Create new instance for usecase weather
func NewNewsUC(tRepo tR.Repository, nRepo nR.Repository) Usecase {
	return &newsUC{
		twitterRepo: tRepo,
		newsRepo:    nRepo,
	}
}
