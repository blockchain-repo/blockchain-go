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

func CreateTable(db string, name string) {
	session := ConnectDB(db)
	respo, err := r.TableCreate(name).RunWrite(session)
	if err != nil {
		log.Printf("Error creating table: %s", err)
	}

	fmt.Printf("%d table created\n", respo.TablesCreated)
}

func CreateDatabase(name string) {
	session := Connect()
	resp, err := r.DBCreate(name).RunWrite(session)
	if err != nil {
		log.Printf("Error creating database: %s", err)
	}

	fmt.Printf("%d DB created\n", resp.DBsCreated)
}

func DropDatabase() {
	dbname := DBNAME
	session := Connect()
	resp, err := r.DBDrop(dbname).RunWrite(session)
	if err != nil {
		log.Printf("Error dropping database: %s", err)
	}

	fmt.Printf("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
}

func InitDatabase() {
	dbname := DBNAME
	CreateDatabase(dbname)

	for _, x := range Tables {
		CreateTable(dbname, x)
	}
}
