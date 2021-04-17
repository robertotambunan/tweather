package weather

import (
	"context"
	"strings"
	"time"
)

func (wu *weatherUC) PostCurrentWeather(ctx context.Context) (err error) {
	var weathers Item

	mapOfCityID := wu.bmkgRepo.GetCityIds(ctx, []string{"Kota Medan", "Kota Jakarta Selatan", "Kab. Tangerang"})

	for cn, mc := range mapOfCityID {
		weatherByID := wu.bmkgRepo.GetWeatherByCityID(ctx, mc)
		if weatherByID.JamCuaca == "" || weatherByID.TemperatureC == "" || weatherByID.Cuaca == "" {
			continue
		}
		weathers.Weathers = append(weathers.Weathers, ItemInfo{
			Place:      strings.Title(cn),
			Weather:    weatherByID.Cuaca,
			Temprature: weatherByID.TemperatureC,
		})
		if weathers.ProjectedTime == "" {
			weathers.ProjectedTime = cleanTime(weatherByID.JamCuaca)
		}
	}

	// To avoid overflow in tweet
	if len(weathers.Weathers) > 3 {
		weathers.Weathers = weathers.Weathers[:3]
	}

	weathers.UpdatedAt = time.Now().Format(time.RFC1123)
	tweet := composeMessage(weathers)
	err = wu.twitterRepo.PostTweet(ctx, tweet)

	return
}

func cleanTime(time string) string {
	return strings.Replace(time, "UTC", "WIB", -1)
}

func composeMessage(weather Item) (tweet string) {
	tweet = "Weather Forcast for " + weather.ProjectedTime + ": "
	for i, w := range weather.Weathers {
		tweet += w.Place + " - "
		tweet += w.Weather + " ("
		tweet += w.Temprature + ")"
		if i < len(weather.Weathers)-1 {
			tweet += ", "
		} else {
			tweet += ". "
		}
	}
	tweet += "This tweet updated at " + weather.UpdatedAt + " by Tweather-Bot."
	return
}
