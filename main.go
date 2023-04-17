package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"github.com/practical-go/go-kyrylo-api/pkg/fetcher"
	"github.com/practical-go/go-kyrylo-api/pkg/input"
)

var newsFetcher fetcher.NewsFetcher

func init() {
	newsFetcher = fetcher.NewNewsProvider(
		fetcher.NewCatFactsNewsFetcher(),
		fetcher.NewSpaceflightNewsFetcher(),
	)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})

	http.HandleFunc("/news", handleNews)

	http.ListenAndServe(":8080", nil)
}

func handleNews(w http.ResponseWriter, r *http.Request) {
	inputReq := input.NewGetNewsRequest(r)

	eTag := fmt.Sprintf(`"%d_%s"`, inputReq.Limit, inputReq.Tag)

	w.Header().Set("Etag", eTag)
	w.Header().Set("Cache-Control", "max-age=60")

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, eTag) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}

	news, err := newsFetcher.GetNews(inputReq)

	if err != nil {
		resp, _ := json.Marshal(domain.ErrorResponse{Error: err.Error()})
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(news)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Write(resp)

	return
}
