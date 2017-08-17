package main

import (
	"os"
	"fmt"

	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/backend"
	"unichain-go/web"
)

func main(){
	fmt.Printf("main:: Hello Unichain-go!\n")
	fmt.Println("main::",os.Args)
	cmd(os.Args)
}

func cmd(args []string) {
	argsCount := len(args)
	if argsCount == 1{
		runHelp()
		return
	}
	switch args[1] {
	case "start":
		runStart()
	case "help":
		runHelp()
	case "configure":
		runConfigure()
	case "init":
		runInit()
	case "drop":
		runDrop()
	case "show-config":
		runShowConfig()
	case "export-my-pubkey":
		runExportMyPubkey()
	default:
		runHelp()
	}
}

func runConfigure() {
	config.ConfigToFile()
}


func runShowConfig() {
	config.FileToConfig()
	fmt.Println(common.Serialize(config.Config))
}

func runExportMyPubkey() {
	config.FileToConfig()
	fmt.Println(config.Config.Keypair.PublicKey)
}

func runInit()  {
	conn :=backend.GetConnection()
	conn.InitDatabase(backend.DBNAME)
}

func runDrop()  {
	conn :=backend.GetConnection()
	conn.DropDatabase(backend.DBNAME)
}

func runStart()  {
	web.Server()
}

func runHelp() {
	fmt.Printf("Commands:\n"+
	"  {configure,show-config,init,drop,start,export-my-pubkey}\n"+
	"	configure           Prepare the config file and create the node keypair\n"+
	"	show-config         Show the current configuration\n"+
	"	export-my-pubkey    Export this node's public key\n"+
	"	init                Init the database\n"+
	"	drop                Drop the database\n"+
	"	start               Start unichain-go\n")
}
