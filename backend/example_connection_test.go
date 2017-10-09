package backend

import (
	"testing"
	"unichain-go/backend/rethinkdb"
	"reflect"
	"unichain-go/log"
)

func ExampleConnection() {
	conn := GetConnection()
	conn.DropDatabase("unichain")
	conn.InitDatabase("unichain")

	//int_res := conn.WriteTransactionToBacklog(`{"id":"5556","back":"j22222ihhh"}`)
	//fmt.Println(int_res)
	//map_string := conn.GetTransactionFromBacklog("5556")
	//fmt.Printf("tx:%s\n", map_string)

	// Output:
	//1 DB dropped, 7 tables dropped
	//1 DB created
	//1 table created
	//1 table created
	//1 table created
	//1 table created
	//1 table created
	//1 table created
	//1 table created
	//{0 1 0 0 0 0 0 0 0 0 0 0 0 0 []  [] []}1
	//tx:{"back":"j22222ihhh","id":"5556"}
}

func Test_createIndex(t *testing.T) {
	conn := GetConnection()
	conn.CreateSecondaryIndex()
}

func Test_testConnection(t *testing.T) {
	//r := rethinkdb.RethinkDBConnection{}
	ty := reflect.TypeOf(rethinkdb.RethinkDBConnection{})
	log.Info(ty)
	s := reflect.New(ty)
	log.Info(s)
}
