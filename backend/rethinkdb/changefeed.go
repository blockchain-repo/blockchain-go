package rethinkdb

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
)

func (c *RethinkDBConnection)Changefeed(db string, name string) *r.Cursor {
	res, err := r.DB(db).Table(name).Changes().Run(c.Session)
	if err != nil {
		log.Print(err)
	}
	return res
}