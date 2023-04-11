package main

import (
	"encoding/json"
	"fmt"
	"github.com/practical-go/go-kyrylo-api/pkg/domain"
	"github.com/practical-go/go-kyrylo-api/pkg/fetcher"
	"net/http"
)

func handleNews(w http.ResponseWriter, _ *http.Request) {
	c := fetcher.NewNewsProvider([]fetcher.NewsFetcher{
		fetcher.NewCatFactsNewsFetcher(),
		fetcher.NewSpaceflightNewsFetcher(),
	})
	news, err := c.GetNews()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	if err != nil {
		resp, _ := json.Marshal(domain.ErrorResponse{Error: err.Error()})
		w.Write(resp)
		return
	}

	resp, _ := json.Marshal(news)
	w.Write(resp)

	return
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})

	http.HandleFunc("/news", handleNews)

	http.ListenAndServe(":8080", nil)
}
