package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
//	"strconv"
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

func InfectNode(node *Node) {
	node.state = I
}

func PrintNode(node *Node) {
	states := [3]string  {"S", "I", "R"}
	fmt.Printf("%d: (%s)\n", node.id, states[node.state])
}

func SIgossip(node *Node, network *[]*Node, m *sync.Mutex, wg *sync.WaitGroup, cycles int, /*c chan string,*/ push bool) {

	for i:=0;i<cycles;i++ {
		// wait

		currState := node.state
		time.Sleep(10 * time.Millisecond)

		// choose a random peer that isn't itself
		var peer *Node = (*network)[rand.Intn(len(*network))]
		for peer == node {
			peer = (*network)[rand.Intn(len(*network))]
		}

//		if(currState == I && peer.state == I) {
//			c <- "Found an infected node, will not send"
//		}

		if(push && currState == I) {
			m.Lock()
				peer.state = I
			m.Unlock()
		}
		time.Sleep(10 * time.Millisecond)
	}
	wg.Done()
}

/*func printer(c chan string) {
	for {
		message := <- c
		fmt.Println(message)
	}
}*/

func main() {
	// For batch testing
	TESTS := 100
	INF := 0
	NETWORK_SIZE := 100
	CYCLES := 6

	for r:=0; r<TESTS; r++ {
		rand.Seed(time.Now().UnixNano())
		var m sync.Mutex
		var wg sync.WaitGroup

//		c := make(chan string)

		// Create a network of NETWORK_SIZE nodes
		network := []*Node{}

		for i := 0; i < NETWORK_SIZE; i++ {
			var x = Node{id: i}
			network = append(network, &x)
		}

		// Infect a randomly selected node from the network
		InfectNode(network[rand.Intn(len(network))])

//		go printer(c) // uncomment for errorcheckign
		for _, node := range network {
			wg.Add(1)
			go SIgossip(node, &network, &m, &wg, CYCLES, /*c,*/ true)
		}
		wg.Wait()

		// Stats
		count := 0
		for _, node := range network {
			if node.state == I {
				count++
			}
		}
		INF = INF + count

		fmt.Printf("%d/%d infected on %d cycles\n", count, NETWORK_SIZE, CYCLES)
	}

	avg := float64(INF) / float64(NETWORK_SIZE)
	pct := avg / float64(25000)
	fmt.Printf("Avg infected: %0.2f (%0.2f)\n", avg, pct)
}
