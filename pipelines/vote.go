package pipelines

import (
	"encoding/json"
	"sync"

	"unichain-go/backend"
	"unichain-go/core"
	"unichain-go/log"
	"unichain-go/models"

	mp "github.com/altairlee/multipipelines/multipipes"
	"unichain-go/common"
)

//TODO init lastVotedBlockId
var lastVotedBlockId string = "ididid-lastVotedBlockId"
var counters map[string]int = make(map[string]int)
var blocksValidityStatus map[string]bool = make(map[string]bool)
var validateChan chan map[string]interface{} = make(chan map[string]interface{})

func validateBlock(arg interface{}) interface{} {
	log.Info("step1: validateBlock")
	blockByte := []byte(arg.(string))
	block := models.Block{}
	err := json.Unmarshal(blockByte, &block)
	//TODO generate the dummy_tx
	if err != nil {
		log.Error(err)
		return []interface{}{block.Id, []string{}}
	}
	if block.ValidateBlock() == false {
		return []interface{}{block.Id, []string{}}
	}
	return []interface{}{block.Id, block.BlockBody.Transactions}
}

func validateTxsInBlock(arg interface{}) interface{} {
	log.Info("step2: validateTxsInBlock")
	blockId := arg.([]interface{})[0].(string)
	txs := arg.([]interface{})[1].([]models.Transaction)
	//doing Parallel start
	txNum := len(txs)
	for index, _ := range txs {
		go func(i int, vc chan map[string]interface{}, bi string, tn int) {
			isValidate := core.ValidateTransaction(txs[i])
			validMap := map[string]interface{}{
				"isValidTx": isValidate,
				"blockId":   bi,
				"txsNum":    tn,
			}
			validateChan <- validMap
		}(index, validateChan, blockId, txNum)
	}

	return blockId
}

func validateBlockByTxs(arg interface{}) interface{} {
	log.Info("step3: validateBlockByTxs")
	//blockId := arg.(string)
	for {
		log.Info(counters)
		validMap := <-validateChan
		isValidTx := validMap["isValidTx"].(bool)
		blockId := validMap["blockId"].(string)
		txNum := validMap["txsNum"].(int)

		counters[blockId] += 1
		isValidBlock, ok := blocksValidityStatus[blockId]
		if !ok {
			isValidBlock = true
		}

		blocksValidityStatus[blockId] = isValidTx && isValidBlock

		if counters[blockId] == txNum {
			return []interface{}{lastVotedBlockId, blockId, blocksValidityStatus[blockId]}
		}
	}

}

func vote(arg interface{}) interface{} {
	log.Info("step4: vote")
	lastVotedBlockId := arg.([]interface{})[0].(string)
	blockId := arg.([]interface{})[1].(string)
	valid := arg.([]interface{})[2].(bool)
	log.Info(lastVotedBlockId)

	//TODO update using lastVotedBlockId
	vote := core.CreateVote(valid, blockId)
	log.Info("Vote `", vote.VoteBody.IsValid, "` for", vote.VoteBody.VoteBlock)

	lastVotedBlockId = blockId
	delete(counters, blockId)
	delete(blocksValidityStatus, blockId)
	return vote
}

func writeVote(arg interface{}) interface{} {
	log.Info("step5: writeVote")
	core.WriteVote(common.Serialize(arg))
	return nil
}

func createVotePipe() (p mp.Pipeline) {
	nodeSlice := make([]*mp.Node, 0)
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateBlock, RoutineNum: 1, Name: "validateBlock"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateTxsInBlock, RoutineNum: 1, Name: "validateBlockTx"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateBlockByTxs, RoutineNum: 1, Name: "validateBlockByTxs"})
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
