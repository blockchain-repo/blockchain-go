package backend

import (
	"unichain-go/backend/rethinkdb"
//	"unichain-go/backend/mongodb"
)

var regStruct map[string]Backend

type Backend interface {
	GetTransaction(id string) map[string]interface{}
	SetTransaction(transaction string) int
}

func init() {
	regStruct = make(map[string]Backend)
	regStruct["rethinkdb"] = &rethinkdb.RethinkDBConnection{}
	//	regStruct["mongodb"] = &mongodb.MongoDBConnection{}
}

func GetBackend() Backend{
	var bd Backend
	str := "rethinkdb"//	TODO Config
	bd = regStruct[str]
	return bd
}