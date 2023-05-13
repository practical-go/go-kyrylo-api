package main

import (
	"encoding/json"
	"fmt"
)

type SpaceflightNewsClient struct {
}

func NewSpaceflightNewsClient() NewsFetcher {
	return &SpaceflightNewsClient{}
}

func (f *SpaceflightNewsClient) GetNews(request *GetNewsRequest) ([]*News, error) {
	body, err := doGetRequest(fmt.Sprintf("https://api.spaceflightnewsapi.net/v4/articles/?limit=%d", request.Limit+1))
	if err != nil {
		return nil, err
	}

	spaceFlightNewsResult := struct {
		Results []struct {
			Title   string `json:"title"`
			Summary string `json:"summary"`
		} `json:"results"`
	}{}

	err = json.Unmarshal(body, &spaceFlightNewsResult)
	if err != nil {
		return nil, err
	}

	var news []*News
	for _, spaceFlightNews := range spaceFlightNewsResult.Results {
		news = append(news, &News{Title: spaceFlightNews.Title, Summary: spaceFlightNews.Summary})
	}

	return news, nil
}
