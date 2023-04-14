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

	var catFactsNewsResult []struct {
		Text string `json:"text"`
	}

	err = json.Unmarshal(body, &catFactsNewsResult)
	if err != nil {
		return nil, err
	}

	news := []*domain.News{}
	for _, catFact := range catFactsNewsResult {
		if catFact.Text == "" {
			continue
		}
		news = append(news, &domain.News{Title: "Cat Facts 😼", Summary: catFact.Text})
	}

	return news, nil
}
