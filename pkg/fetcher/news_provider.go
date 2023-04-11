package fetcher

import (
	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"io/ioutil"
	"math/rand"
	"net/http"
)

type NewsFetcher interface {
	GetNews() ([]*domain.News, error)
}

type NewsProvider struct {
	fetchers []NewsFetcher
}

func NewNewsProvider(fetchers []NewsFetcher) NewsFetcher {
	return &NewsProvider{fetchers: fetchers}
}

func (n *NewsProvider) GetNews() ([]*domain.News, error) {
	result := []*domain.News{}
	for _, fetcher := range n.fetchers {
		news, err := fetcher.GetNews()
		if err != nil {
			return nil, err
		}

		result = append(result, news...)
	}

	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})

	return result, nil
}

func doGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
