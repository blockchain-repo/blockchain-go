package main

import (
	"fmt"

	"unichain-go/common"
//	"unichain-go/config"
	"unichain-go/backend"

//	mp "github.com/altairlee/multipipelines/multipipes"
)

func main(){
	fmt.Printf("Hello Unichain-go!\n")
	//config.ConfigToFile()
	//config.FileToConfig()
	//fmt.Println(common.Serialize(config.Config))

	conn := backend.GetConnection()
	conn.InitDatabase("unichain")
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
	//hash
	fmt.Println(c.Hash("jihao"))
	//generate keypair
	fmt.Println(c.GenerateKeypair("6hXsHQ4fdWQ9UY1XkBYCYRouAagRW8rXxYSLgpveQNYY"))
	msg := "hello unichain 2017"
	pub := "3FyHdZVX4adfSSTg7rZDPMzqzM8k5fkpu43vbRLvEXLJ"
	pub2 := "AZfjdKxEr9G3NwdAkco22nN8PfgQvCr5TDPK1tqsGZrk"
	pri := "5Pv7F7g9BvNDEMdb8HV5aLHpNTNkxVpNqnLTQ58Z5heC"
	sig := "48cpAsUuNf6qKCMFFKitSNjaA8nfPM4o7MacVp8U3QVMbVUr34SSRTTpahi3WEv3GaF2bVWG7J4SLTojgDoacLxT"
	//sign
	sig2 := c.Sign(pri, msg)
	fmt.Println(sig,sig2)
	//verify
	fmt.Println(c.Verify(pub,msg,sig))
	fmt.Println(c.Verify(pub2,msg,sig))

}
