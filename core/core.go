package core

import (
	"unichain-go/models"
	"fmt"
	"unichain-go/common"
)

type Chain struct {
	PublicKey string
	PrivateKey string
	Keyring []string
}

func (c *Chain)CreateTransactionForTest(){
	preOut :=models.PreOut{
		Tx:    "0",
		Index: "0",
	}
	input := models.Input{
		OwnersBefore: c.PublicKey,
		Signature:    "",//TODO
		PreOut:       preOut,
	}
	output :=models.Output{
		OwnersAfter: c.PublicKey,
		Amount:      "1",
	}
	m := map[string]interface{}{}
	m["timestamp"]= common.GenTimestamp()
	tx :=models.Transaction{
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
	fmt.Println(tx.ToString())
}