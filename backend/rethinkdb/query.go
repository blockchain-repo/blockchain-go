package rethinkdb

import (
	"fmt"

	"unichain-go/common"
	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)

func (c *RethinkDBConnection) Get(db string, table string, id string) *r.Cursor {
	res, err := r.DB(db).Table(table).Get(id).Run(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection) Insert(db string, table string, jsonstr string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Insert(r.JSON(jsonstr)).RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection) Update(db string, table string, id string, jsonstr string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Get(id).Update(r.JSON(jsonstr)).RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection) Delete(db string, table string, id string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Get(id).Delete().RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection) GetTransactionFromBacklog(id string) string {
	res := c.Get(DBUNICHAIN, TABLEBACKLOG, id)
	var value map[string]interface{}
	err := res.One(&value)
	map_string := common.Serialize(value)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	return map_string
}

func (c *RethinkDBConnection) WriteTransactionToBacklog(transaction string) int {
	res := c.Insert(DBUNICHAIN, TABLEBACKLOG, transaction)
	return res.Inserted
}

func (c *RethinkDBConnection) DeleteTransaction(id string) int {
	res := c.Delete(DBUNICHAIN, TABLEBACKLOG, id)
	return res.Deleted
}

func (c *RethinkDBConnection) WriteBlock(block string) int {
	res := c.Insert(DBUNICHAIN, TABLEBLOCKS, block)
	return res.Inserted
}

func (c *RethinkDBConnection) WriteVote(vote string) int {
	res := c.Insert(DBUNICHAIN, TABLEVOTES, vote)
	return res.Inserted
}
func (c *RethinkDBConnection) GetUnvotedBlock(pubkey string) []string {
	//TODO doing unfinished lizhen
	res, err := r.DB(DBUNICHAIN).Table(TABLEBLOCKS).Filter(
		func() {

		},
	).Run(c.Session)

	//.Run(c.Session)
	var value []map[string]interface{}
	err = res.All(&value)
	if err != nil {

	}
	//return common.Serialize(value)
	return nil
}

func (c *RethinkDBConnection) GetBlockCount() (int, error) {
	res, err := r.DB(DBUNICHAIN).Table(TABLEBLOCKS).Count().Run(c.Session)
	if err != nil {
		log.Error(err)
		return -1, err
	}
	var cnt int
	res.One(&cnt)
	return cnt, err
}
