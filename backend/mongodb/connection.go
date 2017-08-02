package mongodb

import "log"

type MongoDBConnection struct {

}

func (c *MongoDBConnection)Connect() {
	log.Print("2")
}