package core

import (
	"fmt"
	"testing"

	"unichain-go/common"
	"unichain-go/log"
	"unichain-go/models"
)

func TestCreateBlock(t *testing.T) {
	CreateGenesisBlock()
}

func TestGetBlock(t *testing.T) {
	block := GetBlock("e8e2d19229812d7181bef19aff54741a2219b99447492b79876667a196521089")
	fmt.Println(common.Serialize(block))
}

func TestGetBlocksStatusContainingTx(t *testing.T) {
	log.Debug(GetBlocksStatusContainingTx("c3d2354db940d01446c9088e16066efa1dc16e2a422d42038a4453de6f02ceb5"))
}

func TestIsNewTransaction(t *testing.T) {
	log.Debug(IsNewTransaction("c3d2354db940d01446c9088e16066efa1dc16e2a422d42038a4453de6f02ceb5", ""))
}

func Test_GetLastVotedBlock(t *testing.T) {

	vote := CreateVote(true, "yyy", "")
	WriteVote(common.Serialize(vote))
	res := GetLastVotedBlockId()

	log.Debug(res)
}

func Test_Election(t *testing.T) {

	Election("hhh")

}

func TestBlockElection(t *testing.T) {
	blockId := "eee"
	var votes []models.Vote
	vote1 := CreateVote(true, blockId, "ddd")
	vote2 := CreateVote(false, blockId, "ddd")
	votes = append(votes, vote1)
	votes = append(votes, vote2)
	keyring := []string{PrivateKey}
	log.Debug(BlockElection(blockId, votes, keyring))
}

//func Test_create(t *testing.T) {
//	var txSigners []string = []string{"5XAJvuRGb8B3hUesjREL7zdZ82ahZqHuBV6ttf3UEhyL"}
//	var amount float64 = 100
//	var recipients []interface{} = []interface{}{[]interface{}{"EcWbt741xS8ytvKWEqCPtDu29sgJ1iHubHyoVvuAgc8W", amount}}
//
//	m := map[string]interface{}{}
//	m["timestamp"] = common.GenTimestamp()
//	var metadata map[string]interface{} = m
//	var asset string = "hashid_assert"
//	var chainType string = "unichain"
//	var version string = "1"
//	var relation string = ""
//	var contract string = ""
//
//	tx, err := Create(txSigners, recipients, CREATE, metadata, asset, chainType, version, relation, contract)
//	if err != nil {
//		log.Info(err)
//	}
//	tx.Sign()
//	tx.GenerateId()
//	txMap, err := common.StructToMap(tx)
//	if err != nil {
//		log.Info(err)
//	}
//	WriteTransactionToBacklog(txMap)
//}
