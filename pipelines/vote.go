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

var lastVotedBlockId string = ""
var counters map[string]int = make(map[string]int)
var blocksValidityStatus map[string]bool = make(map[string]bool)

var validateChan = make(chan interface{}, 4000)

func validateBlock(arg interface{}) interface{} {
	log.Debug("step1 validateBlock")
	blockByte := []byte(arg.(string))
	block := models.Block{}
	err := json.Unmarshal(blockByte, &block)
	//TODO generate the dummy_tx  lizhen [method done]
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
	log.Debug("step2 validateTxsInBlock")
	blockId := arg.([]interface{})[0].(string)
	txs := arg.([]interface{})[1].([]models.Transaction)

	txNum := len(txs)
	for index, _ := range txs {
		go func(i int, bi string, tn int) {
			isValidate := core.ValidateTransaction(txs[i])
			validMap := map[string]interface{}{
				"isValidTx": isValidate,
				"blockId":   bi,
				"txsNum":    tn,
			}
			validateChan <- validMap //validMap is the result to the next node
		}(index, blockId, txNum)
	}
	return nil
}

/*
 */
func validateBlockByTxs(arg interface{}) interface{} {
	log.Debug("step3 validateBlockByTxs")

	dataMap := arg.(map[string]interface{})
	isValidTx := dataMap["isValidTx"].(bool)
	chanBlockId := dataMap["blockId"].(string)
	chanTxNum := dataMap["txsNum"].(int)
	counters[chanBlockId] += 1
	isValidBlock, ok := blocksValidityStatus[chanBlockId]
	if !ok {
		isValidBlock = true
	}
	blocksValidityStatus[chanBlockId] = isValidTx && isValidBlock

	if counters[chanBlockId] == chanTxNum {
		return []interface{}{chanBlockId, blocksValidityStatus[chanBlockId]}
	}
	return nil
}

func vote(arg interface{}) interface{} {
	log.Debug("step4 vote")
	blockId := arg.([]interface{})[0].(string)
	valid := arg.([]interface{})[1].(bool)

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
	//TODO get unvoted block lizhen
	return unvotedBlock
}

func createVotePipe() (p mp.Pipeline) {
	nodeSlice := make([]*mp.Node, 0)
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateBlock, RoutineNum: 1, Name: "validateBlock"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateTxsInBlock, RoutineNum: 1, Name: "validateTxsInBlock"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateBlockByTxs, Input: validateChan, RoutineNum: 1, Name: "validateBlockByTxs"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: vote, RoutineNum: 1, Name: "vote"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: writeVote, RoutineNum: 1, Name: "writeVote"})
	p = mp.Pipeline{
		Nodes: nodeSlice,
	}
	return p
}

func getVoteChangefeed() *mp.Node {
	lastVotedBlockId = core.GetLastVotedBlockId()
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
