package pipelines

import (
	"sync"

	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
)


func checkForQuorum(arg interface{}) interface{} {
	return ""
}

func requeueTransactions(arg interface{}) interface{} {
	return ""
}

func createElectionPipe() (p mp.Pipeline) {
	cvNodeSlice := make([]*mp.Node, 0)
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: checkForQuorum, RoutineNum: 1, Name: "checkForQuorum"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: requeueTransactions, RoutineNum: 1, Name: "requeueTransactions"})
	p = mp.Pipeline{
		Nodes: cvNodeSlice,
	}
	return p
}

func getElectionChangefeed() mp.Node {
	conn :=backend.GetConnection()
	node := mp.Node{
		Output:conn.Changefeed("unichain","vote",backend.INSERT),
	}
	return node
}

func StartElectionPipe() {
	p := createElectionPipe()
	changefeed := getElectionChangefeed()
	p.Setup(&changefeed,nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}