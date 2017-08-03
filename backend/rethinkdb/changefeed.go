package rethinkdb

import (
	"unichain-go/log"

	r "gopkg.in/gorethink/gorethink.v3"
)

const (
	INSERT = 1
	DELETE = 2
	UPDATE = 4
)



func (c *RethinkDBConnection)ChangefeedRunForever(operation int){
	var value interface{}
	res := c.GetChangefeed("test", "test")
	for res.Next(&value){
		m := value.(map[string]interface{})
		isInsert := (m["old_val"] == nil)
		isDelete := (m["new_val"] == nil)
		isUpdate := !isInsert && !isDelete
		if isInsert && ((operation & INSERT) != 0) {
			log.Error(m["new_val"])
		}
		if isDelete && ((operation & DELETE) != 0) {
			log.Error(m["old_val"])
		}
		if isUpdate && ((operation & UPDATE) != 0) {
			log.Error(m["new_val"])
		}
	}
}

func (c *RethinkDBConnection)GetChangefeed(db string, table string) *r.Cursor {
	res, err := r.DB(db).Table(table).Changes().Run(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

