package data

import (
	"html/template"
	"time"
)

type Article struct {
	ID string

	CategoryID string

	Title string
	Order int

	Content     string
	ContentHTML template.HTML

	CreatedAt  time.Time
	ModifiedAt time.Time
}
