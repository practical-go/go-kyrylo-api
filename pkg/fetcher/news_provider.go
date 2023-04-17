package fetcher

import (
	"io/ioutil"
	"net/http"

	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"github.com/practical-go/go-kyrylo-api/pkg/input"
)

type NewsFetcher interface {
	GetNews(*input.GetNewsRequest) ([]*domain.News, error)
}

type NewsProvider struct {
	catFactsNewsFetcher    NewsFetcher
	spaceflightNewsFetcher NewsFetcher
}

func NewNewsProvider(catFactsNewsFetcher NewsFetcher, spaceflightNewsFetcher NewsFetcher) NewsFetcher {
	return &NewsProvider{catFactsNewsFetcher, spaceflightNewsFetcher}
}

func (n *NewsProvider) GetNews(req *input.GetNewsRequest) ([]*domain.News, error) {
	result := []*domain.News{}

	var spaceNews []*domain.News
	var catNews []*domain.News
	var err error

	if req.Tag == input.SpaceflightNewsTag {
		spaceNews, err = n.spaceflightNewsFetcher.GetNews(req)
		if err != nil {
			return nil, err
		}
	}

	if req.Tag == input.CatFactsTag {
		catNews, err = n.catFactsNewsFetcher.GetNews(req)
		if err != nil {
			return nil, err
		}
	}

	if req.Tag == input.BothTag {
		catNews, err = n.catFactsNewsFetcher.GetNews(req)
		if err != nil {
			return nil, err
		}

		spaceNews, err = n.spaceflightNewsFetcher.GetNews(req)
		if err != nil {
			return nil, err
		}
	}

	for i, sn, cn := 1, 0, 0; len(result) <= req.Limit; i++ {
		if i%3 > 0 && len(spaceNews) > sn {
			result = append(result, spaceNews[sn])
			sn++
		} else if i%3 == 0 && len(catNews) > cn {
			result = append(result, catNews[cn])
			cn++
		}
	}

	return result[0:req.Limit], nil
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
