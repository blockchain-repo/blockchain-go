package pipelines

import (
	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
)

//changefeed node
type changeNode struct {
	node      mp.Node
	prefeed   []string
	db        string
	table     string
	operation int
}

func (cn *changeNode) runForever() {
	for _, value := range cn.prefeed {
		cn.node.Output <- value
	}
	for {
		//TODO deal error and panic
		cn.getChange()
	}

}

func (cn *changeNode) getChange() {
	conn := backend.GetConnection()
	for i := range conn.Changefeed(cn.db, cn.table, cn.operation) {
		cn.node.Output <- i
	}
}
