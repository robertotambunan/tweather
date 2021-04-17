package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	bR "github.com/robertotambunan/tweather/repositories/bmkg"
)

func (br *bmkgHTTPRepo) GetCityIds(ctx context.Context, cityNames []string) (mapOfCityID map[string]string) {
	req, err := http.NewRequest("GET", wilayahURL, nil)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := br.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var wilayahResp []bR.Wilayah
	json.Unmarshal(bodyBytes, &wilayahResp)

	mapOfCityID = make(map[string]string)

	for _, cn := range cityNames {
		mapOfCityID[strings.ToLower(cn)] = "0"
	}

	for _, wr := range wilayahResp {
		if v, ok := mapOfCityID[strings.ToLower(wr.City)]; ok && v == "0" {
			mapOfCityID[strings.ToLower(wr.City)] = wr.ID
		}
	}

	return
}

func (br *bmkgHTTPRepo) GetWeatherByCityID(ctx context.Context, cityID string) (weather bR.Weather) {
	weatherURLByCity := fmt.Sprintf(weatherURL, cityID)
	req, err := http.NewRequest("GET", weatherURLByCity, nil)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := br.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var weatherResp []bR.Weather
	json.Unmarshal(bodyBytes, &weatherResp)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(loc)

	for _, wr := range weatherResp {

		timeLayout := "2006-01-02 15:04:05"
		tParsed, err := time.Parse(timeLayout, wr.JamCuaca)
		if err != nil {
			continue
		}
		if timeNow.Before(tParsed.Add(-7 * time.Hour)) {
			weather = bR.Weather{
				Cuaca:        wr.Cuaca,
				TemperatureC: wr.TemperatureC + "C",
				JamCuaca:     tParsed.Format(time.RFC1123),
			}
			break
		}
	}
	return
}
