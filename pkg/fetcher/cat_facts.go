package fetcher

import (
	"encoding/json"
	"github.com/practical-go/go-kyrylo-api/pkg/domain"
)

type CatFactsNewsFetcher struct {
	url string
}

func NewCatFactsNewsFetcher() NewsFetcher {
	return &CatFactsNewsFetcher{url: "https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=10"}
}

func (f *CatFactsNewsFetcher) GetNews() ([]*domain.News, error) {
	body, err := doGetRequest(f.url)
	if err != nil {
		return nil, err
	}

	type catFactsNews struct {
		Text string `json:"text"`
	}
	result := []catFactsNews{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	news := []*domain.News{}
	for _, obj := range result {
		if obj.Text == "" {
			continue
		}
		news = append(news, domain.NewNews("Cat Facts 😼", obj.Text))
	}

	return news, nil
}
