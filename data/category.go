package data

import "time"

type Category struct {
	ID string

	Title string
	Order int

	CreatedAt  time.Time
	ModifiedAt time.Time
}
