package rethinkdb

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"
	"fmt"
)


func Get(db string, table string, id string) *r.Cursor {
	session := ConnectDB(db)
	res, err := r.Table(table).Get(id).Run(session)
	if err != nil {
		log.Print(err)
	}
	return res
}

func Insert(db string, table string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(table).Insert(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		log.Print(err)
	}
	return res
}

func Update(db string, table string, id string, jsonstr string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(table).Get(id).Update(r.JSON(jsonstr)).RunWrite(session)
	if err != nil {
		log.Print(err)
	}
	return res
}

func Delete(db string, table string, id string) r.WriteResponse {
	session := ConnectDB(db)
	res, err := r.Table(table).Get(id).Delete().RunWrite(session)
	if err != nil {
		log.Print(err)
	}
	return res
}

func (c *RethinkDBConnection)GetTransaction(id string) map[string]interface{} {
	res := Get("test","test",id)//TODO
	var blo map[string]interface{}
	err := res.One(&blo)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
	}
	return blo
}

func (c *RethinkDBConnection)SetTransaction(transaction string) int {
	res := Insert("test","test",transaction)//TODO
	fmt.Print(res)
	return res.Inserted
}