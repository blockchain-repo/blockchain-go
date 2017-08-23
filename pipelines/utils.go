package pipelines

import (
	"unichain-go/backend"

	mp "github.com/altairlee/multipipelines/multipipes"
)

//changefeed node
type changeNode struct {
	node mp.Node
}

func (cn *changeNode) getChange(db string, table string, operation int) {
	conn := backend.GetConnection()
	for i := range conn.Changefeed(db, table, operation) {
		cn.node.Output <- i
	}
}
