package main

import (
	"go-load-balancer/pkg/balancer"
	"go-load-balancer/pkg/node"
	"go-load-balancer/pkg/task"
	"time"
)

func main() {
	const numNodes = 5

	nodes := make([]*node.Node, numNodes)

	for i := 0; i < numNodes; i++ {
		nodes[i] = node.New(i)
		go nodes[i].Start()
	}

	lb := balancer.NewLeastConnections(nodes)

	const numTasksForSim = 100

	for i := 0; i < numTasksForSim; i++ {
		t := task.New(i, time.Duration(500)*time.Millisecond)
		lb.Distribute(t)
		time.Sleep(time.Millisecond * 100)
	}
}
