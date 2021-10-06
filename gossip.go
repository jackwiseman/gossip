package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
	"strconv"
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

func SIgossip(node *Node, network *[]*Node, m *sync.Mutex, wg *sync.WaitGroup, cycles int, c chan string, push bool) {
	for i:=0;i<cycles;i++ {

		// wait
		time.Sleep(time.Second)

		// choose a random peer
		// need to ensure the node doesn't pick itself
		var peer *Node = (*network)[rand.Intn(len(*network))]
		for peer == node {
			peer = (*network)[rand.Intn(len(*network))]
		}
		// fmt.Printf("%d picked %d\n", node.id, peer.id)

		if(push && node.state == I) {
		c <- strconv.Itoa(node.id) + " -> " + strconv.Itoa(peer.id) + " (cycle " + strconv.Itoa(i+1) + ")"
			m.Lock()
				peer.state = I
			m.Unlock()
		}
	}
	wg.Done()
}

func reader(c chan string) {
	for {
		message := <- c
		fmt.Println(message)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	NETWORK_SIZE := 1000
	CYCLES := 2
	var m sync.Mutex
	var wg sync.WaitGroup


	c := make(chan string)


	network := []*Node{}

	for i := 0; i < NETWORK_SIZE; i++ {
		var x = Node{id: i}
		network = append(network, &x)
	}

	// Update a randomly selected node from the network
	UpdateNode(network[rand.Intn(len(network))])
	go reader(c)
	for _, node := range network {
		wg.Add(1)
		go SIgossip(node, &network, &m, &wg, CYCLES, c, true)
	}
//	wg.Add(1)
//	go func() {
//		wg.Wait()
//		close(c)
//	}()
//	fmt.Println(len(c))
	wg.Wait()

	count := 0
	for _, node := range network {
		if node.state == I {
			count++
		}
	}


	fmt.Printf("%d/%d infected on %d cycles\n", count, NETWORK_SIZE, CYCLES)

}
