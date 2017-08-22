package core

import (
	"math/rand"
	"time"

	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/models"
)

type Chain struct {
}

var PublicKey string
var PrivateKey string
var Keyring []string
var AllPub []string
var Conn backend.Connection

func init() {
	PublicKey = config.Config.Keypair.PublicKey
	PrivateKey = config.Config.Keypair.PrivateKey
	Keyring = config.Config.Keyring
	AllPub = append(Keyring, PublicKey)
	Conn = backend.GetConnection()
}

//Just for test
func CreateTransactionForTest() string {
	preOut := models.PreOut{
		Tx:    "0",
		Index: "0",
	}
	input := models.Input{
		OwnersBefore: PublicKey,
		Signature:    "",
		PreOut:       preOut,
	}
	output := models.Output{
		OwnersAfter: PublicKey,
		Amount:      "1",
	}
	m := map[string]interface{}{}
	m["timestamp"] = common.GenTimestamp()
	tx := models.Transaction{
		Id:        "",
		Inputs:    []models.Input{input},
		Outputs:   []models.Output{output},
		Operation: "CREATE",
		Asset:     "0",
		Chain:     "0",
		Metadata:  m,
		Version:   "1",
	}
	tx.Sign()
	tx.GenerateId()
	return tx.ToString()
}

func InsertToBacklog(m map[string]interface{}) {
	rand.Seed(time.Now().UnixNano())
	//add key
	m["Assign"] = AllPub[rand.Intn(len(AllPub))]
	m["AssignTime"] = common.GenTimestamp()
	str := common.Serialize(m)
	Conn.SetTransactionToBacklog(str)
}

func ValidateTransaction(tx models.Transaction) bool {
	//TODO
	//check hash
	//check sig
	//check asset
	//check input
	//check amoumt
	return true
}
