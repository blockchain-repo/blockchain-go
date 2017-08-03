package rethinkdb

import (
	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func (c *RethinkDBConnection)Changefeed(db string, table string) *r.Cursor {
	res, err := r.DB(db).Table(table).Changes().Run(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}