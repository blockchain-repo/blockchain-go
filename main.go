package main

import (
	"fmt"

	"unichain-go/models"
	"unichain-go/common"
	"unichain-go/log"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")
	tx := models.Transaction{}
	fmt.Println(common.Serialize(tx))
	m,err := common.StructToMap(tx)
	if err != nil {
		log.Error(err.Error())
	}
	delete(m, "id")
	fmt.Println(common.Serialize(m))

}
