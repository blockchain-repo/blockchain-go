package pipelines

import (
	"encoding/json"
	"sync"

	"unichain-go/backend"
	"unichain-go/common"
	"unichain-go/core"
	"unichain-go/log"
	"unichain-go/models"

	mp "github.com/altairlee/multipipelines/multipipes"
)

var txs []models.Transaction

/*
	Filter a transaction.
	Args:
		tx(string)
	Returns:
		The transaction map if assigned to the current node, ``Nil`` otherwise.
*/
func filterTx(arg interface{}) interface{} {
	var m map[string]interface{}
	json.Unmarshal([]byte(arg.(string)), &m)
	if m["Assign"] == core.PublicKey {
		delete(m, "Assign")
		delete(m, "AssignTime")
		return m
	}
	return nil
}

/*
	Validate a transaction.
		Also checks if the transaction already exists in the blockchain.
		If it does, or it's invalid, it's deleted from the backlog immediately.
	Args:
		tx : the transaction to validate.
	Returns:
		:class:`Transaction`: The transaction if valid,
		``None`` otherwise.
*/
func validateTx(arg interface{}) interface{} {
	txByte, err := json.Marshal(arg)
	tx := models.Transaction{}
	err = json.Unmarshal(txByte, &tx)
	if err != nil {
		return nil
	}
	//check already exists
	if core.IsNewTransaction(tx.Id,"") == false {
		core.DeleteTransaction(tx.Id)
		return nil
	}
	//check tx
	if core.ValidateTransaction(tx) == false {
		core.DeleteTransaction(tx.Id)
		return nil
	}
	return tx
}

/*
	"Create a block.
		This method accumulates transactions to put in a block and outputs a block when one of the following conditions is true:
			the size limit of the block has been reached, or a timeout happened.
		Args:
			tx (:class:`Transaction`): the transaction to validate, might be None if a timeout happens.
			timeout (bool): ``True`` if a timeout happened
			(Default: ``False``).
		Returns:
			:class:`Block`: The block, if a block is ready, or ``None``.
*/
func createBlock(arg interface{}) interface{} {
	txs = append(txs, arg.(models.Transaction))
	//TODO when to create [length,timeout] timeout lizhen
	if len(txs) == 1000 || (len(txs) == 2) {
		block := core.CreateBlock(txs)
		var newTxs []models.Transaction
		txs = newTxs
		log.Info("Create Block", block.Id)
		return block
	}
	return nil
}

/*
	Write the block to the Database.
	Args:
		block (:class:`Block`): the block of transactions to write to the database.
	Returns:
		:class:`Block`: The Block.
*/
func writeBlock(arg interface{}) interface{} {
	block := arg.(models.Block)
	core.WriteBlock(common.Serialize(block))
	return block
}

func deleteTransaction(arg interface{}) interface{} {
	block := arg.(models.Block)
	for _,tx := range block.BlockBody.Transactions {
		core.DeleteTransaction(tx.Id)
	}
	return nil
}

func createBlockPipe() (p mp.Pipeline) {
	nodeSlice := make([]*mp.Node, 0)
	nodeSlice = append(nodeSlice, &mp.Node{Target: filterTx, RoutineNum: 1, Name: "filterTx"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: validateTx, RoutineNum: 1, Name: "validateTx"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: createBlock, RoutineNum: 1, Name: "createBlock"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: writeBlock, RoutineNum: 1, Name: "writeBlock"})
	nodeSlice = append(nodeSlice, &mp.Node{Target: deleteTransaction, RoutineNum: 1, Name: "deleteTransaction"})
	p = mp.Pipeline{
		Nodes: nodeSlice,
	}
	return p
}

func getBlockChangeNode() *mp.Node {
	cn := &changeNode{db: "unichain", table: "backlog", operation: backend.INSERT|backend.UPDATE}
	go cn.runForever()
	return &cn.node
}

func StartBlockPipe() {
	log.Info("Block Pipeline Start")
	p := createBlockPipe()
	changeNode := getBlockChangeNode()
	p.Setup(changeNode, nil)
	p.Start()

	waitRoutine := sync.WaitGroup{}
	waitRoutine.Add(1)
	waitRoutine.Wait()
}
