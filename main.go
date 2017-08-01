package main

import (
	"fmt"
	"unichain-go/common"

)

func main(){
	fmt.Printf("Hello Unichain-go!\n")
	fmt.Println(common.GenTimestamp())
	fmt.Println(common.GenDate())
	fmt.Println(common.GenerateUUID())
}