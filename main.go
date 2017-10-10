package main

import (
	"fmt"
	"os"
	"time"

	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/core"
	"unichain-go/log"
	"unichain-go/pipelines"
	"unichain-go/web"
)

func main() {
	args := append(os.Args, "start")
	cmd(args)
}

func cmd(args []string) {
	argsCount := len(args)
	if argsCount == 1 {
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

//First
func runConfigure() {
	config.ConfigToFile()
}

func runShowConfig() {
	fmt.Println(common.Serialize(config.Config))
}

func runExportMyPubkey() {
	fmt.Println(config.Config.Keypair.PublicKey)
}

//Second
func runInit() {
	conn := backend.GetConnection()
	conn.InitDatabase(backend.DBNAME)
	core.CreateGenesisBlock()
	log.Info("Done, have fun!")
}

func runDrop() {
	conn := backend.GetConnection()
	conn.DropDatabase(backend.DBNAME)
}

func runStart() {
	go pipelines.StartBlockPipe()
	//go pipelines.StartVotePipe()
	//go pipelines.StartElectionPipe()
	//go pipelines.StartStalePipe()
	time.Sleep(time.Second * 1)
	web.Server()
}

func runHelp() {
	fmt.Printf("Commands:\n" +
		"  {configure,show-config,init,drop,start,export-my-pubkey}\n" +
		"	configure           Prepare the config file and create the node keypair\n" +
		"	show-config         Show the current configuration\n" +
		"	export-my-pubkey    Export this node's public key\n" +
		"	init                Init the database\n" +
		"	drop                Drop the database\n" +
		"	start               Start unichain-go\n")
}
