package core

import (
	"math/rand"
	"time"

	"errors"
	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/log"
	"unichain-go/models"
)

const (
	CREATE   = "CREATE"
	TRANSFER = "TRANSFER"
	GENESIS  = "GENESIS"
	INTERIM  = "INTERIM"
	CONTRACT = "CONTRACT"
	METADATA = "METADATA"
)
const (
	VERSIONCHAIN    = "1"
	VERSIONCONTRACT = "2" //?
)

type Chain struct {
}

var ALLOWED_OPERATIONS []string = []string{CREATE, TRANSFER, GENESIS, CONTRACT, INTERIM, METADATA}

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

//
func CreateGenesisBlock() string {
	blockCount, err := Conn.GetBlockCount()
	if err != nil {
		log.Error(err)
		return ""
	}
	if blockCount != 0 {
		log.Error("Cannot create the Genesis block!")
		return ""
	}
	block := prepareGenesisBlock()
	log.Info("GenesisBlock: Hello World from the Unichain!")
	WriteBlock(block)
	return block
}

func prepareGenesisBlock() string {
	var txSigners []string = []string{PublicKey}
	var amount int = 1
	var recipients []interface{} = []interface{}{[]interface{}{PublicKey, amount}}
	m := map[string]interface{}{}
	m["message"] = "Hello World from the Unichain"
	var version string = VERSIONCHAIN
	tx, err := CreateTransaction(txSigners, recipients, GENESIS, m, "", "", version, "", "")
	if err != nil {
		log.Info(err)
	}
	tx.Sign()
	tx.GenerateId()
	txs := []models.Transaction{tx}
	block := CreateBlock(txs)
	return block.ToString()
}

func CreateTransaction(txSigners []string, recipients []interface{}, operation string, metadata map[string]interface{}, asset string, chainType string, version string, relation string, contract string) (models.Transaction, error) {
	var tx models.Transaction
	var err error
	if len(txSigners) == 0 {
		err = errors.New("txSigners can not be empty")
		return tx, err
	}
	if len(recipients) == 0 {
		err = errors.New("recipients can not be empty")
		return tx, err
	}
	//TODO do some params validte

	// generate outputs
	var outputs []models.Output
	for _, value := range recipients {
		ownerAfterInfo := value.([]interface{})
		ownerAfter := ownerAfterInfo[0].(string)
		amount := ownerAfterInfo[1].(int)
		output := models.Output{
			OwnersAfter: ownerAfter,
			Amount:      amount,
		}
		outputs = append(outputs, output)
	}

	// generate inputs. Operation CREATE tx only need one and nil preout
	var inputs []models.Input
	var input models.Input = models.Input{
		OwnersBefore: PublicKey,
		Signature:    "",
		PreOut:       nil,
	}
	inputs = append(inputs, input)
	tx = models.Transaction{
		Inputs:    inputs,
		Outputs:   outputs,
		Operation: operation,
		Asset:     asset,
		Chain:     chainType,
		Metadata:  metadata,
		Version:   version,
	}
	return tx, nil
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
		PreOut:       &preOut,
	}
	output := models.Output{
		OwnersAfter: PublicKey,
		Amount:      1,
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

func IsNewTransaction(id string) bool {
	return true
}

func DeleteTransaction(id string) {
	Conn.DeleteTransaction(id)
}

func CreateBlock(txs []models.Transaction) models.Block {
	blockBody := models.BlockBody{
		Transactions: txs,
		NodePubkey:   PublicKey,
		Voters:       AllPub,
		Timestamp:    common.GenTimestamp(),
	}
	block := models.Block{
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

func CreateVote(valid bool, blockId string, previousBlock string) models.Vote {
	voteBody := models.VoteBody{
		IsValid:       valid,
		InvalidReason: "",
		PreviousBlock: previousBlock,
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

func GetUnvotedBlock() []string {
	//TODO get unvoted block
	Conn.GetUnvotedBlock(PublicKey)
	return nil
}
