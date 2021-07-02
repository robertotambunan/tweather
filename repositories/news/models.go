package news

// News : represent news model
type News struct {
	URL   string
	Title string
}

// NewsResp : news api response struct
type NewsResp struct {
	Articles []Article `json:"articles"`
}

// Article : news api response struct for article
type Article struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
