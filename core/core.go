package core

import (
	"math/rand"
	"time"

	"encoding/json"
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
const (
	BLOCK_INVALID = "invalid"
	//return if a block has been voted invalid

	BLOCK_VALID = "valid"
	TX_VALID    = "valid"
	//return if a block is valid, or tx is in valid block

	BLOCK_UNDECIDED = "undecided"
	TX_UNDECIDED    = "undecided"
	//return if block is undecided, or tx is in undecided block

	TX_IN_BACKLOG = "backlog"
	//return if transaction is in backlog
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
	tx.Sign([]string{PrivateKey})
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
	//TODO do some params validate lizhen

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

func CreateDummyTransaction() string {
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
	tx.GenerateId()
	return tx.ToString()
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
	tx.Sign([]string{PrivateKey})
	tx.GenerateId()
	return tx.ToString()
}

func WriteTransactionToBacklog(tx models.Transaction) {
	var m map[string]interface{}
	m, err := common.StructToMap(tx)
	if err != nil {
		log.Error("StructToMap failed")
	}
	rand.Seed(time.Now().UnixNano())
	//add key
	m["Assign"] = AllPub[rand.Intn(len(AllPub))]
	m["AssignTime"] = common.GenTimestamp()
	str := common.Serialize(m)
	Conn.WriteTransactionToBacklog(str)
}

func ValidateTransaction(tx models.Transaction) bool {
	//check hash
	flag := tx.CheckId()
	log.Debug("CheckId tx", tx.Id, flag)
	if flag == false {
		return false
	}
	//check sig
	flag = tx.Verify()
	log.Debug("Verify tx", tx.Id, flag)
	if flag == false {
		return false
	}
	//TODO transfer and others
	//check asset
	//check input
	//check amount
	return true
}

/*
	Return True if the transaction does not exist in any
	VALID or UNDECIDED block. Return False otherwise.
	Args:
	txid (str): Transaction ID
	exclude_block_id (str): Exclude block from search
*/
func IsNewTransaction(id string, exclude_block_id string) bool {
	block_statuses := GetBlocksStatusContainingTx(id)
	delete(block_statuses, exclude_block_id)
	for _, status := range block_statuses {
		if status != BLOCK_INVALID {
			return false
		}
	}
	return true
}

func GetBlocksStatusContainingTx(id string) map[string]string {
	//TODO
	var result map[string]string
	return result
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
	block.Sign(PrivateKey)
	block.GenerateId()
	return block
}

func WriteBlock(block string) {
	Conn.WriteBlock(block)
}

func GetBlock(id string) models.Block {
	var block models.Block
	blockStr := Conn.GetBlock(id)
	err := json.Unmarshal([]byte(blockStr), &block)
	if err != nil {
		log.Error(err)
	}
	return block
}

func ValidateBlock(block models.Block) bool {
	/*
		Validate the Block without validating the transactions.
	*/

	//node_pubkey
	keySet := common.NewHashSet()
	for _, key := range AllPub {
		keySet.Add(key)
	}
	flag := keySet.Has(block.BlockBody.NodePubkey)
	log.Debug("Check node_pubkey", flag)
	if flag == false {
		return false
	}
	//hash
	flag = block.CheckId()
	log.Debug("Check block id", flag)
	if flag == false {
		return false
	}
	//sig
	flag = block.Verify()
	log.Debug("Check block sig", flag)
	if flag == false {
		return false
	}
	//Check that the block contains no duplicated transactions
	txSet := common.NewHashSet()
	for _, tx := range block.BlockBody.Transactions {
		txSet.Add(tx.Id)
	}
	flag = txSet.Len() == len(block.BlockBody.Transactions)
	log.Debug("Check block duplicated transactions", flag, txSet.Len(), len(block.BlockBody.Transactions))
	if flag == false {
		return false
	}
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

func GetUnvotedBlock() []string {
	//TODO get unvoted block lizhen
	//NOT necessary see bigchaindb #1325
	Conn.GetUnvotedBlock(PublicKey)
	return nil
}
