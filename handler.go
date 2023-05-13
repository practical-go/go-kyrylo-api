package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type GetNewsRequest struct {
	Limit int
	Tag   string
}

func NewGetNewsRequest(r *http.Request) *GetNewsRequest {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 0 || limit > 50 {
		limit = 10
	}

	return &GetNewsRequest{Limit: limit, Tag: r.URL.Query().Get("tag")}
}

func HandleNews(newsFetcher NewsFetcher) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		inputReq := NewGetNewsRequest(r)

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
			resp, _ := json.Marshal(ErrorResponse{Error: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(resp)
			return
		}
	
		resp, _ := json.Marshal(news)
	
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Write(resp)
	}

}
