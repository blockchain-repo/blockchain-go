package pipelines

import (
	"sync"

	mp "github.com/altairlee/multipipelines/multipipes"
)


func checkTransactions(arg interface{}) interface{} {
	return ""
}

func reassignTransactions(arg interface{}) interface{} {
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

func StartStalePipe() {
	p := createStalePipe()
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}