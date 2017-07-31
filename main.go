package main

import (
	"fmt"
	"unichain-go/common"
	"unichain-go/core"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")
	fmt.Println(common.GenTimestamp())
	fmt.Println(core.GenDate())
}