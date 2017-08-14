package main

import (
	"fmt"
	"unichain-go/models"
	"unichain-go/common"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")
	fmt.Println(common.Serialize(models.Block{}))
	fmt.Println(common.Serialize(models.Transaction{}))
	fmt.Println(common.Serialize(models.Vote{}))
}
