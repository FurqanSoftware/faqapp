package data

import "time"

type Account struct {
	ID string

	Handle   string
	Password AccountPassword

	FirstIP  string
	RecentIP string

	CreatedAt  time.Time
	ModifiedAt time.Time
}
