package data

import (
	"bytes"
	"crypto/sha1"
	"math/rand"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

type AccountPassword struct {
	Algorithm  string `bson:"algorithm"`
	Salt       []byte `bson:"salt"`
	Iteration  int    `bson:"iteration"`
	KeyLength  int    `bson:"key_length"`
	DerivedKey []byte `bson:"derived_key"`

	CreatedAt time.Time `bson:"created_at"`
}

func NewAccountPassword(clear string) (AccountPassword, error) {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		return AccountPassword{}, err
	}

	pass := AccountPassword{
		Algorithm: "pbkdf2-sha1",
		Salt:      b,
		Iteration: 4096,
		KeyLength: 32,
		CreatedAt: time.Now(),
	}
	pass.DerivedKey = pbkdf2.Key([]byte(clear), pass.Salt, pass.Iteration, pass.KeyLength, sha1.New)
	return pass, nil
}

func (p *AccountPassword) Match(clear string) bool {
	key := pbkdf2.Key([]byte(clear), p.Salt, p.Iteration, p.KeyLength, sha1.New)
	return bytes.Equal(key, p.DerivedKey)
}
