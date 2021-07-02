package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	nR "github.com/robertotambunan/tweather/repositories/news"
)

func (nr *newsHTTPRepo) GetTopNewsBasedOnNation(ctx context.Context, nationCode string, size int) (news []nR.News) {

	if nationCode == "" || size <= 0 {
		return
	}

	req, err := http.NewRequest("GET", topNewsURL, nil)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("apiKey", nr.apiKey)
	q.Add("country", nationCode)
	q.Add("pageSize", strconv.Itoa(size))

	req.URL.RawQuery = q.Encode()

	resp, err := nr.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var newsResp nR.NewsResp

	json.Unmarshal(bodyBytes, &newsResp)

	for _, v := range newsResp.Articles {
		title := v.Title
		if len(title) > 120 {
			title = title[:120] + "..."
		}
		news = append(news, nR.News{
			URL:   v.URL,
			Title: title,
		})
	}

	return
}
