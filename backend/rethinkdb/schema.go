package rethinkdb

import (
	"fmt"

	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)

var Tables = []string{
	"backlog",
	"block",
	"vote",
	"asset",
	"contract",
	"contractvote",
	"contractoutput",
}

func (c *RethinkDBConnection) CreateTable(db string, table string) {
	respo, err := r.DB(db).TableCreate(table).RunWrite(c.Session)
	if err != nil {
		log.Error("Error creating table: %s", err)
	}

	fmt.Printf("%d table created\n", respo.TablesCreated)
}

func (c *RethinkDBConnection) CreateDatabase(db string) {
	resp, err := r.DBCreate(db).RunWrite(c.Session)
	if err != nil {
		log.Error("Error creating database: %s", err)
	}

	fmt.Printf("%d DB created\n", resp.DBsCreated)
}

func (c *RethinkDBConnection) DropDatabase(db string) {
	resp, err := r.DBDrop(db).RunWrite(c.Session)
	if err != nil {
		log.Error("Error dropping database: %s", err)
	}

	fmt.Printf("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
}

func (c *RethinkDBConnection) InitDatabase(db string) {
	c.CreateDatabase(db)
	for _, x := range Tables {
		c.CreateTable(db, x)
	}
}
