package models

type Post struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	PublishedAt string `json:"publishedAt"`
}
