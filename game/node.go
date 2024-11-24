package game

import (
	"blazethenet/text"
	"fmt"
	"go/types"
	"strings"
)

type Node struct {
	Name         text.Text
	LinkedNodes  []Node
	RoutingTable []types.Map
}

func NewNode(name text.Text) *Node {
	return &Node{
		Name: name,
	}
}

func (targetNode *Node) linkNode(linkingNode *Node) {
	targetNode.LinkedNodes = append(targetNode.LinkedNodes, *linkingNode)
}

func (targetNode *Node) String() string {
	var linkedNodeNames []string
	for _, linkedNode := range targetNode.LinkedNodes {
		linkedNodeNames = append(linkedNodeNames, linkedNode.Name.String()) // Assuming text.Text has a String method
	}

	return fmt.Sprintf(
		"Node(Name: %s, LinkedNodes: [%s])",
		targetNode.Name.String(), // Assuming text.Text has a String method
		strings.Join(linkedNodeNames, ", "),
	)
}
