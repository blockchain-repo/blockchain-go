package rethinkdb

import (
	"fmt"

	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)


func (c *RethinkDBConnection)get(db string, table string, id string) *r.Cursor {
	res, err := r.DB(db).Table(table).Get(id).Run(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)insert(db string, table string, jsonstr string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Insert(r.JSON(jsonstr)).RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)update(db string, table string, id string, jsonstr string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Get(id).Update(r.JSON(jsonstr)).RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)delete(db string, table string, id string) r.WriteResponse {
	res, err := r.DB(db).Table(table).Get(id).Delete().RunWrite(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

func (c *RethinkDBConnection)GetTransaction(id string) map[string]interface{} {
	res := c.get("test","test",id)//TODO
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	return blo
}

func (c *RethinkDBConnection)SetTransaction(transaction string) int {
	res := c.insert("test","test",transaction)//TODO
	fmt.Print(res)
	return res.Inserted
}