package rethinkdb

import (
	"testing"
	"unichain-go/log"
	"unichain-go/core"
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

