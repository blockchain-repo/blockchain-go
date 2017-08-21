package pipelines

import (
	"sync"

	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
	"fmt"
)

//Filter a transaction.
//Args:
//tx : the transaction to process.
//Returns:
//dict: The transaction if assigned to the current node,
//``None`` otherwise.
func filterTx(arg interface{}) interface{} {
	fmt.Println("filterTx",arg)
	return ""
}

//Validate a transaction.
//Also checks if the transaction already exists in the blockchain. If it
//does, or it's invalid, it's deleted from the backlog immediately.
//Args:
//tx : the transaction to validate.
//Returns:
//:class:`Transaction`: The transaction if valid,
//``None`` otherwise.
//
func validateTx(arg interface{}) interface{} {
	return ""
}

//"Create a block.
//This method accumulates transactions to put in a block and outputs
//a block when one of the following conditions is true:
//- the size limit of the block has been reached, or
//- a timeout happened.
//Args:
//tx (:class:`Transaction`): the transaction
//to validate, might be None if a timeout happens.
//timeout (bool): ``True`` if a timeout happened
//(Default: ``False``).
//Returns:
//:class:`Block`: The block,
//if a block is ready, or ``None``.
func createBlock(arg interface{}) interface{} {
	return ""
}

//Write the block to the Database.
//Args:
//block (:class:`Block`): the block of
//transactions to write to the database.
//Returns:
//:class:`Block`: The Block.
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
		Output:conn.ChangefeedRunForever("unichain","backlog",backend.INSERT),
	}
	return node
}

func StartBlockPipe() {
	fmt.Println("Block Pipeline Start")
	p := createBlockPipe()
	changefeed := getBlockChangefeed()
	p.Setup(&changefeed,nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}