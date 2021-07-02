package main

import (
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"

	"github.com/robertotambunan/tweather/deliveries"
	bmkgHTTPR "github.com/robertotambunan/tweather/repositories/bmkg/http"
	twHTTPR "github.com/robertotambunan/tweather/repositories/twitter/http"
	nUC "github.com/robertotambunan/tweather/usecases/news"
	wUC "github.com/robertotambunan/tweather/usecases/weather"

	newsHTTPR "github.com/robertotambunan/tweather/repositories/news/http"
)

func main() {
	// load configration
	filePath := "config.toml"
	viper.SetConfigType("toml")
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	logrus.Info("Using Config file: ", viper.ConfigFileUsed())

	if viper.GetString("credential.consumer.key") == "" || viper.GetString("credential.consumer.secret") == "" ||
		viper.GetString("credential.access.token") == "" || viper.GetString("credential.access.secret") == "" {
		log.Fatal("Invalid configuration")
	}

	config := oauth1.NewConfig(viper.GetString("credential.consumer.key"), viper.GetString("credential.consumer.secret"))
	token := oauth1.NewToken(viper.GetString("credential.access.token"), viper.GetString("credential.access.secret"))
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	twClient := twitter.NewClient(httpClient)

	// repositories
	bmkgHTTPRepo := bmkgHTTPR.NewBMKGHTTPRepo(&http.Client{})
	twitterHTTPRepo := twHTTPR.NewTwitterHTTP(twClient)
	newsHTTPRepo := newsHTTPR.NewNewsHTTPRepo(&http.Client{}, viper.GetString("news.api_key"))

	// usecases
	weatherUC := wUC.NewWeatherUC(twitterHTTPRepo, bmkgHTTPRepo)
	newsUC := nUC.NewNewsUC(twitterHTTPRepo, newsHTTPRepo)

	// deliveres
	cronDelivery := deliveries.NewCron(weatherUC, newsUC)

	cronDelivery.ActivateWheater()

}
