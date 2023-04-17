package input

import (
	"net/http"
	"strconv"
)

const (
	SpaceflightNewsTag string = "spaceflight news"
	CatFactsTag               = "cat facts"
	BothTag                   = "both"
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

	tag := r.URL.Query().Get("tag")
	if tag == "" {
		tag = "both"
	}

	return &GetNewsRequest{Limit: limit, Tag: tag}
}
