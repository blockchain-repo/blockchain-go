package multipipes

import (
	"errors"
	"time"
	
	"unichain-go/log"
)

type Node struct {
	Target     func(interface{}) interface{}
	Input      chan interface{}
	Output     chan interface{}
	RoutineNum int
	Name       string
	Timeout    int64
}

func (n *Node) start() {
	for i := 0; i < n.RoutineNum; i++ {
		go n.runForever()
	}
}

func (n *Node) runForever() {
	for {
		//log.Info(n.name, ",in run forever")
		err := n.run()
		if err != nil {
			log.Error(err)
			return
		}
	}
}

func (n *Node) run() error {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * time.Duration(n.Timeout)) //等待10秒钟
		if n.Timeout != 0 {
			timeout <- true
		}
	}()
	select {
	case x, ok := <-n.Input:
		//从ch中读到数据
		if !ok {
			log.Error(errors.New("read data from inputchannel error"))
			return nil
		}
		//TODO  not good enough, how to support multi params and returns
		out := n.Target(x)
		if n.Output == nil || out == nil {
			return nil
		}
		n.Output <- out
	case <-timeout:
		//一直没有从ch中读取到数据，但从timeout中读取到数据
		//log.Info("read data timeout")
		return nil
	}
	return nil
}

type Pipeline struct {
	nodes []*Node
}

func (p *Pipeline) setup(indata *Node) {
	inNode := []*Node{indata}
	nodes_all := append(inNode, p.nodes...)
	p.connect(nodes_all)
}

func (p *Pipeline) connect(nodes []*Node) (ch chan interface{}) {

	if len(nodes) == 0 {
		return nil
	}

	head := nodes[0]
	head.Input = make(chan interface{}, 10)
	head.Output = make(chan interface{}, 10)
	tail := nodes[1:]
	head.Output = p.connect(tail)
	return head.Input
}

func (p *Pipeline) start() {
	for index, _ := range p.nodes {
		p.nodes[index].start()
	}
}