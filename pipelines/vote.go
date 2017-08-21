package pipelines

import (
	"sync"

	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
)


func validateBlock(arg interface{}) interface{} {
	return ""
}


func ungroup(arg interface{}) interface{} {
	return ""
}


func validateBlockTx(arg interface{}) interface{} {
	return ""
}

func vote(arg interface{}) interface{} {
	return ""
}

func writeVote(arg interface{}) interface{} {
	return ""
}

func createVotePipe() (p mp.Pipeline) {
	cvNodeSlice := make([]*mp.Node, 0)
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: validateBlock, RoutineNum: 1, Name: "validateBlock"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: ungroup, RoutineNum: 1, Name: "ungroup"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: validateBlockTx, RoutineNum: 1, Name: "validateBlockTx"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: vote, RoutineNum: 1, Name: "vote"})
	cvNodeSlice = append(cvNodeSlice, &mp.Node{Target: writeVote, RoutineNum: 1, Name: "writeVote"})
	p = mp.Pipeline{
		Nodes: cvNodeSlice,
	}
	return p
}

func getVoteChangefeed() mp.Node {
	conn :=backend.GetConnection()
	node := mp.Node{
		Output:conn.Changefeed("unichain","block",backend.INSERT),
	}
	return node
}

func StartVotePipe() {
	p := createVotePipe()
	changefeed := getVoteChangefeed()
	p.Setup(&changefeed,nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
