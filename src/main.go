package main

import (
	"errors"
	"fmt"
	"heap/heap"
	"os"
)

func main() {
	mp := make(map[rune]int32, 4)
	var nodeSlice []*heap.Node
	var data []byte

	data, err := os.ReadFile("../src/heap/heap-node.go")
	if err != nil {
		panic(errors.New("no such file or directory"))
	}

	for i := 0; i < len(data); i++ {
		if data[i] == '\n' || data[i] == ' ' || data[i] == '\t' {
			continue
		}
		val, in := mp[rune(data[i])]
		if in {
			mp[rune(data[i])] = val + 1
		} else {
			mp[rune(data[i])] = 1
		}
	}

	for k, v := range mp {
		ch := string(rune(k))
		freq := v
		fmt.Println(ch, freq)
		// todo start forming the heap using the first constructor forming the initial leave nodes
		nodeSlice = append(nodeSlice, heap.NewHeapNode(k, v)) // these are currently leaves.
	}

	println()
	for _, node := range nodeSlice {
		fmt.Printf("Node: char=%c, freq=%d\n", node.GetCh(), node.GetFreq())
	}
}
