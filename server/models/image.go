package models

type Image struct {
	Model
	Path        string `json:"path,omitempty"`
	Driver      string `json:"driver,omitempty"`
	Link        string `json:"link,omitempty"`
	Destination string `json:"destination,omitempty"`
	ProjectID   uint64 `json:"project_id,omitempty"`
	PostID      uint64 `json:"post_id,omitempty"`
}
