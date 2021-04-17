package bmkg

import "context"

// Repository : contract for get weather data from bmkg
type Repository interface {
	GetCityIds(ctx context.Context, cityNames []string) (mapOfCityID map[string]string)
	GetWeatherByCityID(ctx context.Context, cityID string) (weather Weather)
}
