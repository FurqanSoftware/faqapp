// Copyright 2018 Furqan Software Ltd. All rights reserved.

package health

import (
	"time"

	"git.furqansoftware.net/faqapp/faqapp/db"
)

func Loop(d time.Duration, dbSess *db.Session) {
	for {
		time.Sleep(d)

		err := dbSess.Ping()
		if err != nil {
			panic("Couldn't reach Mongo")
			return
		}
	}
}
