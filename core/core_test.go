package core

import (
	"fmt"
	"testing"
	"unichain-go/common"
)

func TestCreateBlock(t *testing.T) {
	fmt.Println(PublicKey)
}

func TestGetBlock(t *testing.T) {
	block := GetBlock("e8e2d19229812d7181bef19aff54741a2219b99447492b79876667a196521089")
	fmt.Println(common.Serialize(block))
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
