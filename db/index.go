package db

import "gopkg.in/mgo.v2"

var indexes = map[string][]mgo.Index{
	"articles": {
		{
			Key: []string{"$text:title", "$text:content"},
			Weights: map[string]int{
				"title":   2,
				"content": 1,
			},
		},
	},
}

func MakeIndexes(sess *Session) error {
	for col, indexes := range indexes {
		for _, index := range indexes {
			err := sess.DB("").C(col).EnsureIndex(index)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
