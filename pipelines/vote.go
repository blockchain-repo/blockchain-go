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

func validateBlock(arg interface{}) interface{} {
	blockByte := []byte(arg.(string))
	block := models.Block{}
	err := json.Unmarshal(blockByte, &block)
	//TODO generate the dummy_tx
	if err != nil {
		log.Error(err)
		return []interface{}{block.Id, []string{}}
	}
	if core.ValidateBlock(block) == false {
		return []interface{}{block.Id, []string{}}
	}
	return []interface{}{block.Id, block.BlockBody.Transactions}
}

func validateTxsInBlock(arg interface{}) interface{} {
	blockId := arg.([]interface{})[0].(string)
	txs := arg.([]interface{})[1].([]models.Transaction)
	validateChan := arg.([]interface{})[2].(chan map[string]interface{})
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

	return []interface{}{blockId, txNum, validateChan}
}

/*
 */
func validateBlockByTxs(arg interface{}) interface{} {
	blockId := arg.([]interface{})[0].(string)
	txNum := arg.([]interface{})[1].(int)
	validateChan := arg.([]interface{})[2].(chan map[string]interface{})

	//TODO deal error and panic?
	for i := 0; i < txNum; i++ {
		validMap := <-validateChan
		isValidTx := validMap["isValidTx"].(bool)
		chanBlockId := validMap["blockId"].(string)
		if chanBlockId != blockId {
			log.Error("The func blockId:%s is not the same as the chanBlockId:%s in the channel!", blockId, chanBlockId)
		}
		chanTxNum := validMap["txsNum"].(int)
		if chanTxNum != txNum {
			log.Error("The func txNum:%d is not the same as the chanTxNum:%d in the channel!", txNum, chanTxNum)
		}
		counters[blockId] += 1
		isValidBlock, ok := blocksValidityStatus[blockId]
		if !ok {
			isValidBlock = true
		}
		blocksValidityStatus[blockId] = isValidTx && isValidBlock

	}
	if counters[blockId] == txNum {
		close(validateChan)
		return []interface{}{blockId, blocksValidityStatus[blockId]}
	}
	log.Error("There should be %d txs,but only get %d txs", txNum, counters[blockId])
	return []interface{}{blockId, false}
}

func vote(arg interface{}) interface{} {
	blockId := arg.([]interface{})[0].(string)
	valid := arg.([]interface{})[1].(bool)

	//valid = blocksValidityStatus[blockId]

	vote := core.CreateVote(valid, blockId, lastVotedBlockId)

	log.Info("Vote `", vote.VoteBody.IsValid, "` for", vote.VoteBody.VoteBlock)
	lastVotedBlockId = blockId
	delete(counters, blockId)
	delete(blocksValidityStatus, blockId)
	return vote
}

func writeVote(arg interface{}) interface{} {
	core.WriteVote(common.Serialize(arg))
	return nil
}

func initUnvotedBlock() []string {
	unvotedBlock := []string{}
	//TODO get unvoted block
	return unvotedBlock
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
	preBlock := initUnvotedBlock()
	cn := &changeNode{prefeed: preBlock, db: "unichain", table: "blocks", operation: backend.INSERT}
	go cn.runForever()
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
