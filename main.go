package main

import (
	"fmt"

	"unichain-go/backend"
	"unichain-go/common"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")

	bd := backend.GetBackend()
	map_string := bd.GetTransaction("1111")
	str := common.Serialize(map_string)
	fmt.Printf("tx:%s\n", str)
}