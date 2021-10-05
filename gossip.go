package main

import (
	"fmt"
)

type NodeState int

const (
	S NodeState = iota // Succeptable = 0
	I		   // Infected = 1
	R                  // Removed = 2
)

type Node struct {
	State NodeState
	Neighbors []*Node
	Name string
}

// Connects two nodes for easy initializing of network
func ConnectNodes(n1 *Node, n2 *Node) {
	n1.Neighbors = append(n1.Neighbors, n2)
	n2.Neighbors = append(n2.Neighbors, n1)
}

func PrintNeighbors(node Node) {
	for _, n := range node.Neighbors {
		fmt.Print(n.Name, " ")
	}
}

func main() {
	var A = Node {Name:"A"}
	var B = Node {Name:"B"}
	var C = Node {Name:"C"}
	var D = Node {Name:"D"}
	var E = Node {Name:"E"}
	var F = Node {Name:"F"}
	var G = Node {Name:"G"}
	var H = Node {Name:"H"}
	var I = Node {Name:"I"}
	var J = Node {Name:"J"}
	var K = Node {Name:"K"}

	ConnectNodes(&A, &B)
	ConnectNodes(&B, &C)
	ConnectNodes(&B, &I)
	ConnectNodes(&C, &D)
	ConnectNodes(&C, &F)
	ConnectNodes(&C, &J)
	ConnectNodes(&D, &E)
	ConnectNodes(&D, &K)
	ConnectNodes(&E, &F)
	ConnectNodes(&E, &J)
	ConnectNodes(&F, &G)
	ConnectNodes(&F, &K)
	ConnectNodes(&F, &I)
	ConnectNodes(&G, &J)
	ConnectNodes(&H, &I)
	ConnectNodes(&I, &J)

	PrintNeighbors(F)
}
