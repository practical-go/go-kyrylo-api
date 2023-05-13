package main

import (
	"io/ioutil"
	"net/http"
)

const (
	tagSpace   = "space"
	tagCat     = "cats"
	tagDefault = ""
)

type NewsFetcher interface {
	GetNews(*GetNewsRequest) ([]*News, error)
}

type NewsProvider struct {
	catFactsNewsFetcher    NewsFetcher
	spaceflightNewsFetcher NewsFetcher
}

func NewNewsProvider(catFactsNewsFetcher NewsFetcher, spaceflightNewsFetcher NewsFetcher) NewsFetcher {
	return &NewsProvider{catFactsNewsFetcher, spaceflightNewsFetcher}
}

func (n *NewsProvider) GetNews(req *GetNewsRequest) ([]*News, error) {
	result := []*News{}

	var spaceNews []*News
	var catNews []*News
	var err error

	if req.Tag == tagSpace {
		spaceNews, err = n.spaceflightNewsFetcher.GetNews(req)
		if err != nil {
			return nil, err
		}
	}

	if req.Tag == tagCat {
		catNews, err = n.catFactsNewsFetcher.GetNews(req)
		if err != nil {
			return nil, err
		}
	}

	if req.Tag == tagDefault {
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
