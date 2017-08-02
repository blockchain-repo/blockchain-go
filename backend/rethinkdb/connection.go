package rethinkdb

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

type RethinkDBConnection struct {
	Session *r.Session
}

func (c *RethinkDBConnection)Connect() {
	ip := "localhost"
	session, err := r.Connect(r.ConnectOpts{
		Address: ip + ":28015",
	})

	if err != nil {
		log.Print(err)
	}
	c.Session =session
}