package rethinkdb

import (
	"fmt"
	"unichain-go/common"
	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)


func (c *RethinkDBConnection)Get(db string, table string, id string) *r.Cursor {
	res, err := r.DB(db).Table(table).Get(id).Run(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)Insert(db string, table string, jsonstr string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Insert(r.JSON(jsonstr)).RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)Update(db string, table string, id string, jsonstr string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Get(id).Update(r.JSON(jsonstr)).RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)Delete(db string, table string, id string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Get(id).Delete().RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)GetTransactionFromBacklog(id string) string {
	res := c.Get("unichain","backlog",id)
	var value map[string]interface{}
	err := res.One(&value)
	map_string :=common.Serialize(value)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	return map_string
}

func (c *RethinkDBConnection)SetTransactionToBacklog(transaction string) int {
	res := c.Insert("unichain","backlog",transaction)
	return res.Inserted
}