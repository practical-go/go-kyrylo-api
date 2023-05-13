package main

import (
	"encoding/json"
	"fmt"
)

type CatFactsNewsClient struct {
	url string
}

func NewCatFactsNewsClient() NewsFetcher {
	return &CatFactsNewsClient{}
}

func (f *CatFactsNewsClient) GetNews(request *GetNewsRequest) ([]*News, error) {
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

	var news []*News
	for _, catFact := range catFactsNewsResult {
		if catFact.Text == "" {
			continue
		}
		news = append(news, &News{Title: "Cat Facts 😼", Summary: catFact.Text})
	}

	return news, nil
}
