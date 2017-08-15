package main

import (
	"fmt"

	"unichain-go/models"
	"unichain-go/common"
	"unichain-go/log"
	"unichain-go/config"
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

func runConfigToFile() {
	config.ConfigToFile()
}

func runHelp() {
	fmt.Printf("Commands:\n"+
	"{configure,show-config,init,drop,start}\n"+
	"configure           Prepare the config file and create the node keypair\n"+
	"show-config         Show the current configuration\n"+
	"export-my-pubkey    Export this node's public key\n"+
	"init                Init the database\n"+
	"drop                Drop the database\n"+
	"start               Start unichain-go\n"+
	"command is specific to the MongoDB backend.)\n")
}
