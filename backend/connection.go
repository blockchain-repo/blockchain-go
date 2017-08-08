package backend

import (
	"unichain-go/backend/rethinkdb"
//	"unichain-go/backend/mongodb"


)

var regStruct map[string]Connection

type Connection interface {
	//connection
	Connect()
	//query
	GetTransactionFromBacklog(id string) string
	SetTransactionToBacklog(transaction string) int
	//changefeed TODO
	ChangefeedRunForever(operation int) chan interface{}
	//schema
	InitDatabase(db string)
	DropDatabase(db string)
}

func init() {
	regStruct = make(map[string]Connection)
	regStruct["rethinkdb"] = &rethinkdb.RethinkDBConnection{}
	//regStruct["mongodb"] = &mongodb.MongoDBConnection{}
}

func GetConnection() Connection{
	var conn Connection
	str := "rethinkdb"//TODO Config
	conn = regStruct[str]
	conn.Connect()//needed?
	return conn
}