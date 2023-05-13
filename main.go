package main

import (
	"fmt"
	"net/http"
)

var newsFetcher NewsFetcher

func init() {
	newsFetcher = NewNewsProvider(
		NewCatFactsNewsClient(),
		NewSpaceflightNewsClient(),
	)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, World")
	})

	http.HandleFunc("/news", HandleNews(newsFetcher))

	http.ListenAndServe(":8081", nil)
}
