package pipelines

import (
	"sync"

	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
)

func filterTx(arg interface{}) interface{} {
	return ""
}


func validateTx(arg interface{}) interface{} {
	return ""
}

func createBlock(arg interface{}) interface{} {
	return ""
}

func writeBlock(arg interface{}) interface{} {
	return ""
}

func createBlockPipe() (p mp.Pipeline) {
	cvNodeSlice := make([]*mp.Node, 0)
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: filterTx, RoutineNum: 1, Name: "filterTx"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: validateTx, RoutineNum: 1, Name: "validateTx"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: createBlock, RoutineNum: 1, Name: "createBlock"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: writeBlock, RoutineNum: 1, Name: "writeBlock"})
	p = mp.Pipeline{
		Nodes: cvNodeSlice,
	}
	return p
}

func getBlockChangefeed() mp.Node {
	conn :=backend.GetConnection()
	node := mp.Node{
		Output:conn.ChangefeedRunForever("unichain","backend",1),
	}
	return node
}

func StartBlockPipe() {
	p := createBlockPipe()
	changefeed := getBlockChangefeed()
	p.Setup(&changefeed,nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}