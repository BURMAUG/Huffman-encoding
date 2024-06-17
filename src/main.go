package main

import (
	"errors"
	"fmt"
	"heap/heap"
	"os"
)

var (
	nodeSlice []*heap.Node
	data      []byte
	hp        *heap.Heap
)

func main() {
	hp = &heap.Heap{}
	data, err := os.ReadFile("../src/heap/heap-node.go")
	if err != nil {
		panic(errors.New("no such file or directory"))
	}
	// fmt.Println(string(data))
	mp := makeRuneFrequencyMap(data)

	// Making individual nodes in order to mkae the heap
	for k, v := range mp {
		nodeSlice = append(nodeSlice, heap.NewHeapNode(k, v)) // these are currently leaves.
		hp.Insert(heap.NewHeapNode(k, v))                     // added to the heap. AKA Making the forrest
	}

	// create new nodes using the second constructor of the heap node
	for hp.Size() > 1 {

		firstNode, err := hp.DeleteMin()
		if err != nil {
			panic(errors.New("err from first node"))
		}

		secondNode, err := hp.DeleteMin()
		if err != nil {
			panic(errors.New("err from second node"))
		}

		parent := heap.NewHeapNodes(0, firstNode.GetFreq()+secondNode.GetFreq(), firstNode, secondNode)
		hp.Insert(parent)
	}

	cur := hp.Peek()
	// have to start building the huffman encoding
	for _, node := range nodeSlice {
		fmt.Printf("Node: char=%c, freq=%d\n", node.GetCh(), node.GetFreq())
	}
	fmt.Printf("Node: char=%c, freq=%d\n", cur.GetCh(), cur.GetFreq())

	println()

	huffmanEncoding := make(map[rune]string, 1) // my huffman encoding table.
	preorderEncoding(cur, "", huffmanEncoding)

	for k, v := range huffmanEncoding {
		fmt.Printf("char=%c -> %s\n", k, v)
	}

	println()

	encoding := ""
	for i := 0; i < len(data); i++ {
		encoding += huffmanEncoding[rune(data[i])]
	}
	fmt.Println(encoding) //compressed!
	text, _ := os.Create("compress.txt")
	text.WriteString(encoding)
	text.Close()
}

func preorderEncoding(node *heap.Node, s string, encoding map[rune]string) {

	if node.Left() == nil && node.Right() == nil && node.GetCh() != 0 {
		encoding[node.GetCh()] = s
		// fmt.Printf("Node: char=%c, freq=%s\n", node.GetCh(), s)
		return
	}

	preorderEncoding(node.Left(), s+string("0"), encoding)
	preorderEncoding(node.Right(), s+string("1"), encoding)
}

func makeRuneFrequencyMap(data []byte) map[rune]int {
	// here is where the data has to be re
	mp := make(map[rune]int)

	for i := 0; i < len(data); i++ {
		val, in := mp[rune(data[i])]

		if in {
			mp[rune(data[i])] = val + 1
		} else {
			mp[rune(data[i])] = 1
		}

	}

	return mp
}
