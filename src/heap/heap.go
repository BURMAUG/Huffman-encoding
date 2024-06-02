package heap

type Heap struct {
	nodes []*Node
}

func (h *Heap) Parent(index int) int {
	return index - 1/2
}

// LeftChild
func LeftChild(index int) int {
	return 2*index + 1
}

// RightChild
func RightChild(index int) int {
	return 2*index + 2
}

// HasLeftChild
func (h *Heap) HasLeftChild(index int) bool {
	return LeftChild(index) < len(h.nodes)
}

// HasRightChild
func (h *Heap) HasRightChild(index int) bool {
	return RightChild(index) < len(h.nodes)
}

// Swap
func (h *Heap) Swap(index1, index2 int) {
	h.nodes[index1], h.nodes[index2] = h.nodes[index2], h.nodes[index1]
}

// Peek retrieves but does not remove element from heap
func (h *Heap) Peek() *Node {
	return h.nodes[0]
}

func (h *Heap) DeleteMin() *Node {
	return nil
}

func (h *Heap) Insert(n Node) {

}

func (h *Heap) Size() int32 {
	return 2
}
