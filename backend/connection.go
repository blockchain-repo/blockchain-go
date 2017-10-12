package backend

import (
	"unichain-go/backend/rethinkdb"
	//	"unichain-go/backend/mongodb"
	"unichain-go/config"
)

const (
	INSERT = 1
	DELETE = 2
	UPDATE = 4
	DBNAME = "unichain"
)

var regStruct map[string]Connection

type Connection interface {
	//connection
	Connect()
	//query //table.backlog
	GetTransactionFromBacklog(id string) string
	WriteTransactionToBacklog(transaction string) int
	DeleteTransaction(id string) int
	//query //table.blocks
	GetBlock(id string) string
	WriteBlock(block string) int
	GetBlockCount() (int, error)
	GetUnvotedBlock(pubkey string) []string
	//query //table.votes
	WriteVote(vote string) int
	//changefeed
	Changefeed(db string, table string, operation int) chan interface{}
	//schema
	InitDatabase(db string)
	DropDatabase(db string)
}

func init() {
	regStruct = make(map[string]Connection)
	regStruct["rethinkdb"] = &rethinkdb.RethinkDBConnection{}
	//regStruct["mongodb"] = &mongodb.MongoDBConnection{}
}

func GetConnection() Connection {
	var conn Connection
	str := config.Config.Database.Backend
	conn, ok := regStruct[str]
	if !ok {
		return nil
	}
	conn.Connect() //needed?
	return conn
}
