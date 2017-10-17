package pipelines

import (
	"encoding/json"
	"sync"
	"time"

	"unichain-go/common"
	"unichain-go/core"
	"unichain-go/log"
	"unichain-go/models"

	mp "github.com/altairlee/multipipelines/multipipes"
)

var timeout time.Duration = 5
var reassignDelay time.Duration = 5

func checkTransactions(arg interface{}) interface{} {
	//TODO timeout
	time.Sleep(timeout * time.Second)
	txs := core.GetStaleTransaction(reassignDelay)
	log.Debug(txs)
	return txs
}

func reassignTransactions(arg interface{}) interface{} {
	var txsMap []map[string]interface{}
	json.Unmarshal([]byte(arg.(string)), &txsMap)
	for _, tx := range txsMap {
		delete(tx, "Assign")
		delete(tx, "AssignTime")
		txStr := common.Serialize(tx)
		var txModel models.Transaction
		json.Unmarshal([]byte(txStr), &txModel)
		core.UpdateTransactionToBacklog(txModel)
	}
	return ""
}

func createStalePipe() (p mp.Pipeline) {
	cvNodeSlice := make([]*mp.Node, 0)
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: checkTransactions, RoutineNum: 1, Name: "checkTransactions"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: reassignTransactions, RoutineNum: 1, Name: "reassignTransactions"})
	p = mp.Pipeline{
		Nodes: cvNodeSlice,
	}
	return p
}

type preNode struct {
	node mp.Node
}

func (p *preNode) produceData() {
	//note you can init some datas before start produce
	for {
		s := "produce data"
		p.node.Output <- s
		time.Sleep(timeout * time.Second)
	}
}

func startProduceData() *mp.Node {
	pre := &preNode{}
	go pre.produceData()
	return &pre.node
}
func StartStalePipe() {
	log.Info("Stale Pipeline Start")
	p := createStalePipe()
	inData := startProduceData()
	p.Setup(inData, nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
