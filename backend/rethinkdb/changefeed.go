package rethinkdb

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
)

func Changefeed(db string, name string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(name).Changes().Run(session)
	if err != nil {
		log.Print(err)
	}
	return res
}