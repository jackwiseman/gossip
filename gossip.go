package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type NodeState int

const (
	S NodeState = iota // Succeptable = 0
	I		   // Infected = 1
	R		   // Removed = 2
)

type Node struct {
	state NodeState `default:S`
	id int
}

func UpdateNode(node *Node) {
	node.state = I
}

func PrintNode(node *Node) {
	states := [3]string  {"S", "I", "R"}
	fmt.Printf("%d: (%s)\n", node.id, states[node.state])
}

func SIgossip(node *Node, network *[]*Node, m *sync.Mutex, wg *sync.WaitGroup, push bool) {
	// wait

	// choose a random peer

	// need to ensure the node doesn't pick itself
	var peer *Node = (*network)[rand.Intn(len(*network))]
	for peer == node {
		peer = (*network)[rand.Intn(len(*network))]
	}
	fmt.Printf("%d picked %d\n", node.id, peer.id)
	
	m.Lock()
	if(push && node.state == I) {
		UpdateNode(peer)
	}
	m.Unlock()

	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	NETWORK_SIZE := 5
	var m sync.Mutex
	var wg sync.WaitGroup
	network := []*Node{}

	for i := 0; i < NETWORK_SIZE; i++ {
		var x = Node{id: i}
		network = append(network, &x)
	}

	// Update a randomly selected node from the network
	UpdateNode(network[rand.Intn(len(network))])

	fmt.Println("\n== Before ==")
	for _, node := range network {
		PrintNode(node)
	}

	for _, node := range network {
		wg.Add(1)
		go SIgossip(node, &network, &m, &wg, true)
	}

	wg.Wait()

	fmt.Println("\n== After ==")
	for _, node := range network {
		PrintNode(node)
	}
}
