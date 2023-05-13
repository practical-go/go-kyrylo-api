package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestHandleNews(t *testing.T) {

	req, err := http.NewRequest("GET", "/news?limit=10&tag=foo", nil)
	
	if err != nil {
        t.Fatal(err)
    }

	rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HandleNews(&NewsFetcherMock{}))
    

    handler.ServeHTTP(rr, req)
    
	assert.Equal(t, rr.Code, http.StatusOK)

	assert.JSONEq(t, `[{"title":"foo","summary":"bar"}]`, rr.Body.String())
}

func TestHandleNewsWithErrorResponse(t *testing.T) {

	req, err := http.NewRequest("GET", "/news?limit=10&tag=foo", nil)
	
	if err != nil {
        t.Fatal(err)
    }

	rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HandleNews(&NewsFetcherErrorResponseMock{}))
    

    handler.ServeHTTP(rr, req)
    
	assert.Equal(t, rr.Code, http.StatusInternalServerError)

	assert.JSONEq(t, `{"Error": "Error while fetching news"}`, rr.Body.String())
}

type NewsFetcherMock struct {}

func (*NewsFetcherMock) GetNews(req *GetNewsRequest) ([]*News, error) {
	return []*News{
		{
			Title: "foo",
			Summary:  "bar",
		},
	}, nil
}

type NewsFetcherErrorResponseMock struct {}

func (*NewsFetcherErrorResponseMock) GetNews(req *GetNewsRequest) ([]*News, error) {
	return nil, fmt.Errorf("Error while fetching news")
}