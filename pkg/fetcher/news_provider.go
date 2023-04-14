package fetcher

import (
	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"io/ioutil"
	"net/http"
)

type NewsFetcher interface {
	GetNews() ([]*domain.News, error)
}

type NewsProvider struct {
	catFactsNewsFetcher    NewsFetcher
	spaceflightNewsFetcher NewsFetcher
}

func NewNewsProvider(catFactsNewsFetcher NewsFetcher, spaceflightNewsFetcher NewsFetcher) NewsFetcher {
	return &NewsProvider{catFactsNewsFetcher, spaceflightNewsFetcher}
}

func (n *NewsProvider) GetNews() ([]*domain.News, error) {
	result := []*domain.News{}

	spaceNews, err := n.spaceflightNewsFetcher.GetNews()
	if err != nil {
		return nil, err
	}

	catNews, err := n.catFactsNewsFetcher.GetNews()
	if err != nil {
		return nil, err
	}

	for i := 0; len(result) < 10; i++ {
		if len(spaceNews) > i+2 {
			result = append(result, spaceNews[i:i+2]...)
		}
		if len(catNews) > i {
			result = append(result, catNews[i])
		}
	}

	return result[0:10], nil
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
