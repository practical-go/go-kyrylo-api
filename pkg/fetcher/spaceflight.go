package fetcher

import (
	"encoding/json"
	"github.com/practical-go/go-kyrylo-api/pkg/domain"
)

type SpaceflightNewsFetcher struct {
	url string
}

func NewSpaceflightNewsFetcher() NewsFetcher {
	return &SpaceflightNewsFetcher{url: "https://api.spaceflightnewsapi.net/v4/articles/?limit=10"}
}

func (f *SpaceflightNewsFetcher) GetNews() ([]*domain.News, error) {
	body, err := doGetRequest(f.url)
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

	var news []*domain.News
	for _, spaceFlightNews := range spaceFlightNewsResult.Results {
		news = append(news, &domain.News{Title: spaceFlightNews.Title, Summary: spaceFlightNews.Summary})
	}

	return news, nil
}
