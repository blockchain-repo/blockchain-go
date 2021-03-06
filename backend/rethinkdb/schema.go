package rethinkdb

import (
	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)

const (
	DBUNICHAIN = "unichain"

	TABLEBACKLOG         = "backlog"
	TABLEBLOCKS          = "blocks"
	TABLEVOTES           = "votes"
	TABLEASSETS          = "assets"
	TABLECONTRACTS       = "contracts"
	TABLECONTRACTVOTES   = "contractvotes"
	TABLECONTRACTOUTPUTS = "contractoutputs"
)

var Tables = []string{
	"backlog",
	"blocks",
	"votes",
	"assets",
	"contracts",
	"contractvotes",
	"contractoutputs",
}

func (c *RethinkDBConnection) CreateSecondaryIndex() {
	//Create backlog index

	//Create blocks index
	response, err := r.DB(DBUNICHAIN).Table(TABLEBLOCKS).IndexCreateFunc(
		"transaction_id",
		r.Row.Field("BlockBody").Field("Transactions").Field("id"),
		r.IndexCreateOpts{Multi:true},
	).RunWrite(c.Session)
	if err != nil {
		log.Error("Error creating index:", err)
	}
	log.Info("%d index created", response.Created)
	//Create votes index
	response, err = r.DB(DBUNICHAIN).Table(TABLEVOTES).IndexCreateFunc("block_and_voter",
		func(row r.Term) interface{} {
			return []interface{}{row.Field("VoteBody").Field("VoteBlock"), row.Field("NodePubkey")}
		},
	).RunWrite(c.Session)
	if err != nil {
		log.Error("Error creating index:", err)
	}
	log.Info("%d index created", response.Created)
}

func (c *RethinkDBConnection) CreateTable(db string, table string) {
	respo, err := r.DB(db).TableCreate(table).RunWrite(c.Session)
	if err != nil {
		log.Error("Error creating table:", err)
	}

	log.Info("%d table created %s", respo.TablesCreated, table)
}

func (c *RethinkDBConnection) CreateDatabase(db string) {
	resp, err := r.DBCreate(db).RunWrite(c.Session)
	if err != nil {
		log.Error("Error creating database:", err)
	}

	log.Info("%d DB created %s", resp.DBsCreated, db)
}

func (c *RethinkDBConnection) DropDatabase(db string) {
	resp, err := r.DBDrop(db).RunWrite(c.Session)
	if err != nil {
		log.Error("Error dropping database:", err)
	}

	log.Info("%d DB dropped, %d tables dropped\n", resp.DBsDropped, resp.TablesDropped)
}

func (c *RethinkDBConnection) InitDatabase(db string) {
	c.CreateDatabase(db)
	for _, x := range Tables {
		c.CreateTable(db, x)
	}
	c.CreateSecondaryIndex()
}
