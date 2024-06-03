package main

import (
	"errors"
	"fmt"
	"heap/heap"
	"os"
)

func main() {
	var nodeSlice []*heap.Node
	var data []byte
	hp := &heap.Heap{}

	data, err := os.ReadFile("../src/heap/heap-node.go")
	if err != nil {
		panic(errors.New("no such file or directory"))
	}

	mp := makeRuneFrequencyMap(data)

	// Making individual nodes in order to mkae the heap
	for k, v := range mp {
		nodeSlice = append(nodeSlice, heap.NewHeapNode(k, v)) // these are currently leaves.
		hp.Insert(heap.NewHeapNode(k, v))                     // added to the heap. AKA Making the forrest
	}

	// create new nodes using the second constructor of the heap node
	for i := 0; i < hp.Size() && hp.Size() >= 2; i++ {
		firstNode, err := hp.DeleteMin()
		if err != nil {
			panic(errors.New("err from first node"))
		}
		secondNode, err := hp.DeleteMin()
		if err != nil {
			panic(errors.New("err from second node"))
		}
		hp.Insert(heap.NewHeapNodes(0, firstNode.GetFreq()+secondNode.GetFreq(), nil, nil))
	}

	// have to start building the huffman encoding
	for _, node := range nodeSlice {
		fmt.Printf("Node: char=%c, freq=%d\n", node.GetCh(), node.GetFreq())
	}
}

func makeRuneFrequencyMap(data []byte) map[rune]int {
	// here is where the data has to be re
	mp := make(map[rune]int)
	for i := 0; i < len(data); i++ {
		//	if data[i] == '\n' || data[i] == ' ' || data[i] == '\t' {
		//	continue
		//}
		val, in := mp[rune(data[i])]
		if in {
			mp[rune(data[i])] = val + 1
		} else {
			mp[rune(data[i])] = 1
		}
	}
	return mp
}
