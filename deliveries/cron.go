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

	err = cw.newsUsecase.PostTopNewsIndonesia(context.Background())
	if err != nil {
		log.Println(err)
	}

	err = cw.newsUsecase.PostTopNewsSport(context.Background())
	if err != nil {
		log.Println(err)
	}
}
