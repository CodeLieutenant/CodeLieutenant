package models

type Post struct {
	Model
	Title   string `json:"title,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Content string `json:"content,omitempty"`
}
