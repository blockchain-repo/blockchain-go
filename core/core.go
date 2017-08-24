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

func WriteTransactionToBacklog(m map[string]interface{}) {
	rand.Seed(time.Now().UnixNano())
	//add key
	m["Assign"] = AllPub[rand.Intn(len(AllPub))]
	m["AssignTime"] = common.GenTimestamp()
	str := common.Serialize(m)
	Conn.WriteTransactionToBacklog(str)
}

func ValidateTransaction(tx models.Transaction) bool {
	//TODO
	//check hash
	//check sig
	//check asset
	//check input
	//check amount
	return true
}

func CreateBlock(txs []models.Transaction) models.Block {
	blockBody := models.BlockBody{
		Transactions: txs,
		NodePubkey:   PublicKey,
		Voters:       AllPub,
		Timestamp:    common.GenTimestamp(),
	}
	block := models.Block{
		Id:        "",
		BlockBody: blockBody,
		Signature: "",
	}
	block.Sign()
	block.GenerateId()
	return block
}

func WriteBlock(block string) {
	Conn.WriteBlock(block)
}

func ValidateBlock(block models.Block) bool {
	//TODO
	return true
}

func CreateVote(valid bool, blockId string) models.Vote {
	voteBody := models.VoteBody{
		IsValid:       valid,
		InvalidReason: "",
		//TODO PreviousBlock
		PreviousBlock: "",
		VoteBlock:     blockId,
		Timestamp:     common.GenTimestamp(),
	}
	vote := models.Vote{
		Id:         "",
		NodePubkey: PublicKey,
		VoteBody:   voteBody,
		Signature:  "",
	}
	vote.Sign()
	vote.GenerateId()
	return vote
}

func WriteVote(vote string) {
	Conn.WriteVote(vote)
}

func Election(blockId string) bool {
	//TODO
	return true
}

func Requeue(blockId string) {

}