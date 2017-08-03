package rethinkdb

import (
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
)

func (c *RethinkDBConnection)Changefeed(db string, table string) *r.Cursor {
	res, err := r.DB(db).Table(table).Changes().Run(c.Session)
	if err != nil {
		log.Print(err)
	}
	return res
}