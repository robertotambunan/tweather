package deliveries

import (
	"github.com/robertotambunan/tweather/usecases/news"
	"github.com/robertotambunan/tweather/usecases/weather"
)

// CronWeather delivery for triggered cron
type CronWeather struct {
	weatherUsecase weather.Usecase
	newsUsecase    news.Usecase
}

// NewCron Initiate New cron
func NewCron(weatherUC weather.Usecase, newsUC news.Usecase) *CronWeather {
	return &CronWeather{
		weatherUC,
		newsUC,
	}
}
