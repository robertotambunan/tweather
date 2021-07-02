package http

import (
	"net/http"

	nR "github.com/robertotambunan/tweather/repositories/news"
)

type newsHTTPRepo struct {
	httpClient *http.Client
	apiKey     string
}

const (
	topNewsURL = "https://newsapi.org/v2/top-headlines"
)

// NewNewsHTTPRepo create new instance for bmkg http repo
func NewNewsHTTPRepo(httpClient *http.Client, apiKey string) nR.Repository {
	return &newsHTTPRepo{
		httpClient: httpClient,
		apiKey:     apiKey,
	}
}
