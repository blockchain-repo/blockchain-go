package main

import (
	"fmt"

	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/log"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")

	conn := backend.GetConnection()
	map_string := conn.GetTransaction("1111")
	str := common.Serialize(map_string)
	fmt.Printf("tx:%s\n", str)
	int_res := conn.SetTransaction(`{"back":"j22222ihhh"}`)
	fmt.Println(int_res)
	conn.RunChangeFeed(1)

	log.Error("1")
}
