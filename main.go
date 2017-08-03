package main

import (
	"fmt"

	"unichain-go/backend"
	mp "github.com/altairlee/multipipelines/multipipes"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")

	conn := backend.GetConnection()
	map_string := conn.GetTransaction("1111")
	fmt.Printf("tx:%s\n", map_string)
	int_res := conn.SetTransaction(`{"back":"j22222ihhh"}`)
	fmt.Println(int_res)

	node := mp.Node{
		Output:conn.ChangefeedRunForever(1),
	}
	for i := range node.Output {
		fmt.Println(i)
	}
}
