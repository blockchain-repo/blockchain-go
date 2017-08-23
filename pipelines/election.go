package pipelines

import (
	"sync"

	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
)

func checkForQuorum(arg interface{}) interface{} {
	//TODO
	//core.election(blockId string)
	return ""
}

func requeueTransactions(arg interface{}) interface{} {
	//TODO
	//core.requeue(blockId string)
	return ""
}

func createElectionPipe() (p mp.Pipeline) {
	nodeSlice := make([]*mp.Node, 0)
	nodeSlice = append(nodeSlice, &mp.Node{Target: checkForQuorum, RoutineNum: 1, Name: "checkForQuorum"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: requeueTransactions, RoutineNum: 1, Name: "requeueTransactions"})
	p = mp.Pipeline{
		Nodes: nodeSlice,
	}
	return p
}

func getElectionChangeNode() *mp.Node {
	cn := &changeNode{}
	go cn.getChange("unichain", "vote", backend.INSERT)
	return &cn.node
}

func StartElectionPipe() {
	p := createElectionPipe()
	changeNode := getElectionChangeNode()
	p.Setup(changeNode, nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
