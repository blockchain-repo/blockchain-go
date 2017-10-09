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
