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

func validateBlock(arg interface{}) interface{} {
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

func validateBlockTx(arg interface{}) interface{} {
	blockId := arg.([]interface{})[0].(string)
	txs := arg.([]interface{})[1].([]models.Transaction)
	var valid bool
	//TODO Parallel
	for _, tx := range txs {
		if core.ValidateTransaction(tx) == false {
			valid = false
			return []interface{}{valid, blockId}
		}
	}
	valid = true
	return []interface{}{valid, blockId}
}

func vote(arg interface{}) interface{} {
	valid := arg.([]interface{})[0].(bool)
	blockId := arg.([]interface{})[1].(string)
	vote := core.CreateVote(valid, blockId)
	log.Info("Vote `", vote.VoteBody.IsValid, "` for", vote.VoteBody.VoteBlock)
	return vote
}

func writeVote(arg interface{}) interface{} {
	core.WriteVote(common.Serialize(arg))
	return nil
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
