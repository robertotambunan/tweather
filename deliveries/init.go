package deliveries

import (
	"github.com/robertotambunan/tweather/usecases/weather"
)

// CronWeather delivery for triggered cron
type CronWeather struct {
	weatherUsecase weather.Usecase
}

// NewCron Initiate New cron
func NewCron(weatherUC weather.Usecase) *CronWeather {
	return &CronWeather{
		weatherUC,
	}
}
