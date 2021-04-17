package weather

import "context"

// ItemInfo : represnt weather info that we want to show to user
type ItemInfo struct {
	Place      string
	Weather    string
	Temprature string
}

// Item : collection of weather
type Item struct {
	Weathers      []ItemInfo
	UpdatedAt     string
	ProjectedTime string
}

// Usecase : contract for usecases for weather need
type Usecase interface {
	PostCurrentWeather(ctx context.Context) (err error)
}
