package rethinkdb

import (
	"unichain-go/log"
	"unichain-go/common"

	r "gopkg.in/gorethink/gorethink.v3"
)

const (
	INSERT = 1
	DELETE = 2
	UPDATE = 4
)



func (c *RethinkDBConnection)ChangefeedRunForever(operation int) chan interface{} {
	var value interface{}
	ch := make(chan interface{})
	res := c.GetChangefeed("test", "test")
	go func() {
		for res.Next(&value){
			m := value.(map[string]interface{})
			isInsert := (m["old_val"] == nil)
			isDelete := (m["new_val"] == nil)
			isUpdate := !isInsert && !isDelete
			if isInsert && ((operation & INSERT) != 0) {
				ch<- common.Serialize(m["new_val"])
			}
			if isDelete && ((operation & DELETE) != 0) {
				ch <- common.Serialize(m["old_val"])
			}
			if isUpdate && ((operation & UPDATE) != 0) {
				ch <- common.Serialize(m["new_val"])
			}
		}
	}()
	return ch
}

func (c *RethinkDBConnection)GetChangefeed(db string, table string) *r.Cursor {
	res, err := r.DB(db).Table(table).Changes().Run(c.Session)
	if err != nil {
		log.Error(err)
	}
	return res
}

