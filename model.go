package main

type News struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

type ErrorResponse struct {
	Error string
}

