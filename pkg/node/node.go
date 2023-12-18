package node

import (
	"fmt"
	"go-load-balancer/pkg/task"
	"sync"
	"time"
)

type Node struct {
	ID                int
	ActiveConnections int
	TaskChannel       chan *task.Task
	WG                *sync.WaitGroup
}

func New(id int) *Node {
	return &Node{
		ID:                id,
		ActiveConnections: 0,
		TaskChannel:       make(chan *task.Task),
		WG:                &sync.WaitGroup{},
	}
}

func (n *Node) Start() {
	for t := range n.TaskChannel {
		n.processTask(t)
	}
}

func (n *Node) processTask(t *task.Task) {
	n.WG.Add(1)
	defer n.WG.Done()

	n.ActiveConnections++
	fmt.Printf("Node %d processing task %d\n", n.ID, t.ID)

	time.Sleep(time.Duration(t.ProcessingTime) * time.Nanosecond)

	n.ActiveConnections--
	fmt.Printf("Node %d completed task %d\n", n.ID, t.ID)
}

func (n *Node) SendTask(t *task.Task) {
	n.TaskChannel <- t
}

func (n *Node) GetActiveConnections() int {
	return n.ActiveConnections
}

func (n *Node) Stop() {
	close(n.TaskChannel)
	n.WG.Wait()
}
