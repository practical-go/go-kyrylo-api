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

	type spaceflightResultObject struct {
		Results []struct {
			Title   string `json:"title"`
			Summary string `json:"summary"`
		} `json:"results"`
	}
	result := spaceflightResultObject{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	news := []*domain.News{}
	for _, res := range result.Results {
		news = append(news, domain.NewNews(res.Title, res.Summary))
	}

	return news, nil
}
