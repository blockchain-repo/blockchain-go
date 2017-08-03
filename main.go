package main

import (
	"fmt"

	"unichain-go/backend"
	"unichain-go/log"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")

	conn := backend.GetConnection()
	map_string := conn.GetTransaction("1111")
	fmt.Printf("tx:%s\n", map_string)
	int_res := conn.SetTransaction(`{"back":"j22222ihhh"}`)
	fmt.Println(int_res)

	for i := range conn.ChangefeedRunForever(1).Output{
		fmt.Println(i)
	}

	log.Error("1")
}
