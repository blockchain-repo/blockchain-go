package pipelines

import (
	"encoding/json"
	"fmt"
	"sync"

	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/core"
	"unichain-go/models"

	mp "github.com/altairlee/multipipelines/multipipes"
)

// Filter a transaction.
// Args: tx(string)
// Returns: The transaction map if assigned to the current node, ``Nil`` otherwise.
func filterTx(arg interface{}) interface{} {
	fmt.Println("filterTx", arg)
	var m map[string]interface{}
	json.Unmarshal([]byte(arg.(string)), &m)
	if m["Assign"] == core.PublicKey {
		delete(m, "Assign")
		delete(m, "AssignTime")
		return m
	}
	return nil
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
	fmt.Println("validateTx", common.Serialize(arg))
	//check already exists
	//check tx
	txByte, err := json.Marshal(arg)
	if err != nil {
		return nil
	}
	tx := models.Transaction{}
	err = json.Unmarshal(txByte, &tx)
	if err != nil {
		return nil
	}
	if core.ValidateTransaction(tx) == false {
		return nil
	}
	return tx
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
	fmt.Println("createBlock", common.Serialize(arg))
	var txs []models.Transaction
	txs = append(txs, arg.(models.Transaction))
	flag := true
	//TODO when to create
	if flag == true {
		block := core.CreateBlock(txs)
		return block
	}
	return nil
}

//Write the block to the Database.
//Args:
//block (:class:`Block`): the block of
//transactions to write to the database.
//Returns:
//:class:`Block`: The Block.
func writeBlock(arg interface{}) interface{} {
	fmt.Println("writeBlock", common.Serialize(arg))
	core.WriteBlock(common.Serialize(arg))
	return nil
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

func getBlockChangeNode() *mp.Node {
	cn := &changeNode{}
	go cn.getChange("unichain", "backlog", backend.INSERT)
	return &cn.node
}

func StartBlockPipe() {
	p := createBlockPipe()
	changeNode := getBlockChangeNode()
	p.Setup(changeNode, nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
