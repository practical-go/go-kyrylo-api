package fetcher

import (
	"encoding/json"
	"fmt"

	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"github.com/practical-go/go-kyrylo-api/pkg/input"
)

type CatFactsNewsFetcher struct {
	url string
}

func NewCatFactsNewsFetcher() NewsFetcher {
	return &CatFactsNewsFetcher{}
}

func (f *CatFactsNewsFetcher) GetNews(request *input.GetNewsRequest) ([]*domain.News, error) {
	body, err := doGetRequest(fmt.Sprintf("https://cat-fact.herokuapp.com/facts/random?animal_type=cat&amount=%d", request.Limit+1))
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
