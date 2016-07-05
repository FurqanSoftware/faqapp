package data

import "time"

type AccountPassword struct {
	Algorithm  string
	Salt       []byte
	Iteration  int
	KeyLength  int
	DerivedKey []byte

	CreatedAt time.Time
}
