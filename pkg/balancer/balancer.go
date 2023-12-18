package balancer

import (
	"go-load-balancer/pkg/node"
	"go-load-balancer/pkg/task"
)

type Balancer struct {
	nodes []*node.Node
}

func NewLeastConnections(nodes []*node.Node) *Balancer {
	return &Balancer{
		nodes: nodes,
	}
}

func (b *Balancer) Distribute(t *task.Task) {
	var leastLoadedNode *node.Node
	minConnections := -1

	for _, n := range b.nodes {
		activeConnections := n.GetActiveConnections()
		if minConnections == -1 || activeConnections < minConnections {
			minConnections = activeConnections
			leastLoadedNode = n
		}
	}

	if leastLoadedNode != nil {
		leastLoadedNode.SendTask(t)
	}
}
