package rethinkdb

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

type RethinkDBConnection struct {}

func Connect() *r.Session {
	ip := "localhost"
	session, err := r.Connect(r.ConnectOpts{
		Address: ip + ":28015",
	})

	if err != nil {
		log.Print(err)
	}
	return session
}

func ConnectDB(dbname string) *r.Session {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: dbname,
	})

	if err != nil {
		log.Print(err)
	}
	return session
}