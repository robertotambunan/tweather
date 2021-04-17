package http

import (
	"net/http"

	bR "github.com/robertotambunan/tweather/repositories/bmkg"
)

type bmkgHTTPRepo struct {
	httpClient *http.Client
}

const (
	wilayahURL = "https://ibnux.github.io/BMKG-importer/cuaca/wilayah.json"
	weatherURL = "https://ibnux.github.io/BMKG-importer/cuaca/%s.json"
)

// NewBMKGHTTPRepo create new instance for bmkg http repo
func NewBMKGHTTPRepo(httpClient *http.Client) bR.Repository {
	return &bmkgHTTPRepo{
		httpClient: httpClient,
	}
}
