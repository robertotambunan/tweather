package news

import "context"

// Usecase : contract for usecases for news need
type Usecase interface {
	PostTopNewsIndonesia(ctx context.Context) (err error)
	PostTopNewsSport(ctx context.Context) (err error)
}
