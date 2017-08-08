package main

import (
	"fmt"

	"unichain-go/common"
//	"unichain-go/config"
//	"unichain-go/backend"

//	mp "github.com/altairlee/multipipelines/multipipes"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")
	//config.ConfigToFile()
	//config.FileToConfig()
	//fmt.Println(common.Serialize(config.Config))

//	conn := backend.GetConnection()
////	conn.InitDatabase("unichain")
//	map_string := conn.GetTransactionFromBacklog("1111")
//	fmt.Printf("tx:%s\n", map_string)
//	int_res := conn.SetTransactionToBacklog(`{"back":"j22222ihhh"}`)
//	fmt.Println(int_res)
//
//	node := mp.Node{
//		Output:conn.ChangefeedRunForever(1),
//	}
//	for i := range node.Output {
//		fmt.Println(i)
//	}
	c :=common.GetCrypto()
	fmt.Println(c.Hash("jihao"))
}
