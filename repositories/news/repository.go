package news

import "context"

// Repository : contract for get weather data from bmkg
type Repository interface {
	GetTopNewsBasedOnNation(ctx context.Context, nationCode string, size int) (news []News)
}
