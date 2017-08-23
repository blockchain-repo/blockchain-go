package pipelines

import (
	"encoding/json"
	"sync"

	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/log"
	"unichain-go/models"

	mp "github.com/altairlee/multipipelines/multipipes"
	"fmt"
)

func validateBlock(arg interface{}) interface{} {
	log.Info("step1: validateBlock:", arg)
	var m map[string]interface{}
	json.Unmarshal([]byte(arg.(string)), &m)


	//TODO err
	blockByte := []byte(arg.(string))
	block := models.Block{}
	err := json.Unmarshal(blockByte, &block)
	if err != nil {
		fmt.Printf("%T\n",blockByte)
		log.Error(err)
		return nil
	}
	return nil

	bs, err := json.Marshal(arg)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	log.Info(bs)

	block1 := models.Block{}
	err = json.Unmarshal(bs, &block1)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	log.Info(common.Serialize(block))

	//TODO validate the block is the same as the block in db

	//validate block content
	err = block.ValidateBlock()
	if err != nil {
		//TODO generate the dummy_tx
		return []interface{}{block.Id, []string{}}
	}
	return []interface{}{block.Id, block.BlockBody.Transactions}
}

func validateBlockTx(arg interface{}) interface{} {
	log.Info("step2: validateBlockTx", arg)
	blockId := arg.([]interface{})[0].(string)
	log.Info("blockId:", blockId)
	txs := arg.([]interface{})[1].([]models.Transaction)
	for index, tx := range txs {
		log.Info(index)
		log.Info(common.Serialize(tx))
	}
	return ""
}

func vote(arg interface{}) interface{} {
	return ""
}

func writeVote(arg interface{}) interface{} {
	return ""
}

func createVotePipe() (p mp.Pipeline) {
	nodeSlice := make([]*mp.Node, 0)
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateBlock, RoutineNum: 1, Name: "validateBlock"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateBlockTx, RoutineNum: 1, Name: "validateBlockTx"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: vote, RoutineNum: 1, Name: "vote"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: writeVote, RoutineNum: 1, Name: "writeVote"})
	p = mp.Pipeline{
		Nodes: nodeSlice,
	}
	return p
}

func getVoteChangefeed() *mp.Node {
	cn := &changeNode{}
	go cn.getChange("unichain", "block", backend.INSERT)
	return &cn.node
}

func StartVotePipe() {
	log.Info("Vote Pipeline Start")
	p := createVotePipe()
	changeNode := getVoteChangefeed()
	p.Setup(changeNode, nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
