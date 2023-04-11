package domain

type News struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

func NewNews(title, summary string) *News {
	return &News{Title: title, Summary: summary}
}
