package backend

import (
	"unichain-go/backend/rethinkdb"
	//	"unichain-go/backend/mongodb"
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
	GetTransactionFromBacklog(id string) string
	WriteTransactionToBacklog(transaction string) int

	WriteBlock(block string) int
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
	str := "rethinkdb" //TODO Config
	conn = regStruct[str]
	conn.Connect() //needed?
	return conn
}
