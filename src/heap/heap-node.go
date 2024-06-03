package heap

import "errors"

// nodeInterface is used for inherited features
type nodeInterface interface {
	GetCh() rune
	GetFreq() int
	Left() (*Node, error)
	Right() (*Node, error)
}

// Node represents a node to a heap
type Node struct {
	ch    rune
	freq  int
	left  *Node
	right *Node
}

// NewHeapNode is a way to create a new heap given a character and frequency.
// it returns a pointer to the created node
func NewHeapNode(ch rune, freq int) *Node {
	return &Node{ch: ch, freq: freq}
}

// NewHeapNodes is a way to create a new heap given a character and frequency.
// it returns a pointer to the created node
func NewHeapNodes(ch rune, freq int, left, right *Node) *Node {
	return &Node{ch: ch, freq: freq, left: left, right: right}
}

func (n *Node) GetCh() rune {
	return n.ch
}

func (n *Node) GetFreq() int {
	return n.freq
}

func (n *Node) Left() (*Node, error) {
	if n.left == nil {
		panic(errors.New("left is nil"))
	}
	return n.left, nil
}

func (n *Node) Right() *Node {
	return n.right
}
