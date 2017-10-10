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
	//query
	//transaction
	GetTransactionFromBacklog(id string) string
	WriteTransactionToBacklog(transaction string) int
	DeleteTransaction(id string) int
	//block
	WriteBlock(block string) int
	GetBlockCount() (int, error)
	GetUnvotedBlock(pubkey string) []string
	//vote
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
	str := config.Config.Database.Host
	conn = regStruct[str]
	conn.Connect() //needed?
	return conn
}
