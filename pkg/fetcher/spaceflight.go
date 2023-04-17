package fetcher

import (
	"encoding/json"
	"fmt"

	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"github.com/practical-go/go-kyrylo-api/pkg/input"
)

type SpaceflightNewsFetcher struct {
}

func NewSpaceflightNewsFetcher() NewsFetcher {
	return &SpaceflightNewsFetcher{}
}

func (f *SpaceflightNewsFetcher) GetNews(request *input.GetNewsRequest) ([]*domain.News, error) {
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

	var news []*domain.News
	for _, spaceFlightNews := range spaceFlightNewsResult.Results {
		news = append(news, &domain.News{Title: spaceFlightNews.Title, Summary: spaceFlightNews.Summary})
	}

	return news, nil
}
