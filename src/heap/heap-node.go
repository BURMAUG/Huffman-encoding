package heap

import "errors"

// NodeInterface is used for inherited features
type NodeInterface interface {
	GetCh() rune
	GetFreq() int32
	Left() (*Node, error)
	Right() (*Node, error)
}

// Node represents a node to a heap
type Node struct {
	ch    rune
	freq  int32
	left  *Node
	right *Node
}

// NewHeapNode is a way to create a new heap given a character and frequency.
// it returns a pointer to the created node
func NewHeapNode(ch rune, freq int32) *Node {
	return &Node{ch: ch, freq: freq}
}

func NewHeapNodes(ch rune, freq int32, left, right Node) *Node {
	return &Node{ch: ch, freq: freq, left: &left, right: &right}
}

func (n *Node) GetCh() rune {
	return n.ch
}

func (n *Node) GetFreq() int32 {
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
