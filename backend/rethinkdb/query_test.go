package rethinkdb

import (
	"testing"
	"unichain-go/log"
)

func Test_GetBlockCount(t *testing.T) {
	c := &RethinkDBConnection{}
	c.Connect()

	res, err := c.GetBlockCount()
	log.Info(err)
	log.Info(res)
}

func Test_GetBlock(t *testing.T) {
	c := &RethinkDBConnection{}
	c.Connect()

	res:= c.GetBlock("e8e2d19229812d7181bef19aff54741a2219b99447492b79876667a196521089")
	log.Info(res)
}

func Test_GetGenesisBlock(t *testing.T) {
	c := &RethinkDBConnection{}
	c.Connect()

	res:= c.GetGenesisBlock()
	log.Info(res)
}

func TestRethinkDBConnection_GetBlocksContainTransaction(t *testing.T) {
	c := &RethinkDBConnection{}
	c.Connect()
	log.Debug(c.GetBlocksContainTransaction("c3d2354db940d01446c9088e16066efa1dc16e2a422d42038a4453de6f02ceb5"))
}

func Test_rql(t *testing.T) {
	c := &RethinkDBConnection{}
	c.Connect()
	//.Filter(r.Row.Field("Operation").Eq("GENESIS"))
	//.Filter(r.Row.Field("BlockBody").Field("Transactions"))
	//Eq([]interface{}{map[string]interface{}{"Operation":"GENESIS",}}))
	//r.Row.Field("BlockBody").Field("Transactions").Nth(0).Filter(map[string]interface{}{"Operation":"GENESIS",})
	res, err := r.DB("unichain").Table("blocks").Filter("").Run(c.Session)
	log.Info(err)
	var value map[string]interface{}
	//var key []string
	res.One(&value)
	log.Info(value)
	log.Info(common.Serialize(value))
}
