package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
	"os"
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

func InfectNode(node *Node) {
	node.state = I
}

func PrintNode(node *Node) {
	states := [3]string  {"S", "I", "R"}
	fmt.Printf("%d: (%s)\n", node.id, states[node.state])
}

func SIgossip(node *Node, network *[]*Node, m *sync.Mutex, wg *sync.WaitGroup, cycles int, /*c chan string,*/ push bool, pull bool) {

	for i:=0;i<cycles;i++ {
		// wait

		currState := node.state
		time.Sleep(time.Millisecond)

		// choose a random peer that isn't itself
		var peer *Node = (*network)[rand.Intn(len(*network))]
		for peer == node {
			peer = (*network)[rand.Intn(len(*network))]
		}

		currPeerState := peer.state
		time.Sleep(time.Millisecond)

		if(push && currState == I) {
			m.Lock()
				peer.state = I
			m.Unlock()
		}

		if(pull && currPeerState == I) {
//				c <- strconv.Itoa(node.id) + " requesting from " + strconv.Itoa(peer.id) + "(cycle " + strconv.Itoa(i) + ")"
			m.Lock()
//			updateRequest(peer, node, c)
				node.state = I
			m.Unlock()
		}
		time.Sleep(time.Millisecond)
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
	INF := 0
	var NETWORK_SIZE int
	var CYCLES int
	var TESTS int
	METHOD := "push"
	defaults := [3]int {10, 4, 10} // network_size, cycles, tests

	// Command line input
	if (len(os.Args[1:]) > 1) { 
		if(os.Args[1] == "help") {
			fmt.Println("Usage: gossip [push/pull/push-pull] [network_size] [cycles] [tests]")
			return
		} else {
			METHOD = os.Args[1]
		}
	}

	if (len(os.Args[1:]) > 1) {
		for x:=0; x<len(os.Args[2:]); x++ {
			i, _ := strconv.Atoi(os.Args[x+2])
			defaults[x] = i 
		}
	}

	NETWORK_SIZE = defaults[0]
	CYCLES = defaults[1]
	TESTS = defaults[2]

	fmt.Println("Running " + METHOD + " on network of size " + strconv.Itoa(NETWORK_SIZE) + " with " + strconv.Itoa(CYCLES) + " cycles " + strconv.Itoa(TESTS) + " times...")
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
			switch METHOD {
				case "push":
					go SIgossip(node, &network, &m, &wg, CYCLES, /*c,*/ true, false)
				case "pull":
					go SIgossip(node, &network, &m, &wg, CYCLES, /*c,*/ false, true)
				case "push-pull":
					go SIgossip(node, &network, &m, &wg, CYCLES, /*c,*/ true, true)
			}
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

//		fmt.Printf("%d/%d infected on %d cycles\n", count, NETWORK_SIZE, CYCLES)
//		fmt.Println("")
	}

	avg := float64(INF) / float64(TESTS)
	fmt.Printf("Avg infected: %0.2f \n", avg)
}
