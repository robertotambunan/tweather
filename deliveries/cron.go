package deliveries

import (
	"context"
	"log"
)

// ActivateWheaterUpdateCron function to activate cron for wheater
func (cw *CronWeather) ActivateWheater() {
	err := cw.weatherUsecase.PostCurrentWeather(context.Background())
	if err != nil {
		log.Println(err)
	}
}
