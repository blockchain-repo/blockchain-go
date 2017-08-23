package pipelines

import (
	"encoding/json"
	"testing"
	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/config"
	"unichain-go/log"
	"unichain-go/models"
)

//Just for test
func CreateTransactionForTest() models.Transaction {

	var PublicKey string
	PublicKey = config.Config.Keypair.PublicKey
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
	return tx
}

//Just for test
func CreateBlockForTest() models.Block {
	//PublicKey := config.Config.Keypair.PublicKey
	//txs := make([]models.Transaction, 0)
	//for i := 1; i < 5; i++ {
	//	tx := CreateTransactionForTest()
	//	txs = append(txs, tx)
	//}
	var bloclBody models.BlockBody = models.BlockBody{
	//Transactions: txs,
	//NodePubkey:   PublicKey,
	//Voters:       []string{PublicKey},
	//Timestamp:    common.GenTimestamp(),
	}
	var block models.Block = models.Block{
		BlockBody: bloclBody,
	}
	//block.Sign()
	//block.GenerateId()
	arg := common.Serialize(block)
	bs, err := json.Marshal(arg)
	if err != nil {
		log.Error(err.Error())
		//return nil
	}
	log.Info(bs)

	blockStru := models.Block{}
	err = json.Unmarshal(bs, &blockStru)
	if err != nil {
		log.Error(err.Error())
		//return nil
	}
	return block
}

func TestInsertBlock(t *testing.T) {
	Conn := backend.GetConnection()
	block := CreateBlockForTest()
	Conn.WriteBlock(common.Serialize(block))
}

func TestVotePip(t *testing.T) {
	StartVotePipe()
}
