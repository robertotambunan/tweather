package weather

import (
	bR "github.com/robertotambunan/tweather/repositories/bmkg"
	tR "github.com/robertotambunan/tweather/repositories/twitter"
)

type weatherUC struct {
	twitterRepo tR.Repository
	bmkgRepo    bR.Repository
}

// NewWeatherUC Create new instance for usecase weather
func NewWeatherUC(tRepo tR.Repository, bRepo bR.Repository) Usecase {
	return &weatherUC{
		twitterRepo: tRepo,
		bmkgRepo:    bRepo,
	}
}
