package rethinkdb

import (
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
)

const DBNAME = "test"

var Tables = []string{
	"transaction",
}

func (c *RethinkDBConnection)CreateTable(db string, table string) {
	respo, err := r.DB(db).TableCreate(table).RunWrite(c.Session)
	if err != nil {
		log.Printf("Error creating table: %s", err)
	}

	fmt.Printf("%d table created\n", respo.TablesCreated)
}

func (c *RethinkDBConnection)CreateDatabase(db string) {
	resp, err := r.DBCreate(db).RunWrite(c.Session)
	if err != nil {
		log.Printf("Error creating database: %s", err)
	}

	fmt.Printf("%d DB created\n", resp.DBsCreated)
}

func (c *RethinkDBConnection)DropDatabase(db string) {
	resp, err := r.DBDrop(db).RunWrite(c.Session)
	if err != nil {
		log.Printf("Error dropping database: %s", err)
	}

	fmt.Printf("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
}

func (c *RethinkDBConnection)InitDatabase() {
	dbname := DBNAME
	c.CreateDatabase(dbname)
	for _, x := range Tables {
		c.CreateTable(dbname, x)
	}
}
