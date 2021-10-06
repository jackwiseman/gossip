package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NodeState int

const (
	S NodeState = iota // Succeptable = 0
	I		   // Infected = 1
	R		   // Removed = 2
)

type Node struct {
	state NodeState `default:S`
	//name string
	id int
}

func UpdateNode(node *Node) {
	node.state = I
}

func PrintNode(node *Node) {
	states := [3]string  {"S", "I", "R"}
	fmt.Printf("%d: (%s)\n", node.id, states[node.state])
}

func SIgossip(node *Node, network *[]*Node, push bool) {
	// wait

	// choose a random peer
	x := node.id
	// need to ensure the node doesn't pick itself
	for x == node.id {
		x = rand.Intn(len(*network))
	}
	var peer *Node = (*network)[x]
	fmt.Printf("%d picked %d\n", node.id, peer.id)

	if(push && node.state == I) {
		UpdateNode(peer)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	NETWORK_SIZE := 5

	network := []*Node{}

	for i := 0; i < NETWORK_SIZE; i++ {
		var x = Node{id: i}
		network = append(network, &x)
	}

	// Update a randomly selected node from the network
	UpdateNode(network[rand.Intn(len(network))])

	fmt.Println("== Before ==")
	for _, node := range network {
		PrintNode(node)
	}

	for _, node := range network {
		go SIgossip(node, &network, true)
	}

	fmt.Println("== After ==")
	for _, node := range network {
		PrintNode(node)
	}
}
