package heap

type Heap struct {
	nodes []*Node
}

func Parent(index int) int {
	return (index - 1) / 2
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

// DeleteMin deletes the root node as this is the minimum.
func (h *Heap) DeleteMin() *Node {
	return nil
}

// HeapifyUp Move nodes with the min value to the top of the Heap.
func (h *Heap) HeapifyUp(index int) {
	parentIndex := Parent(index)
	if index > 0 && h.nodes[index].GetFreq() < h.nodes[parentIndex].GetFreq() {
		h.Swap(index, parentIndex)
		h.HeapifyUp(parentIndex)
	}
}

func (h *Heap) HeapifyDown(index int) {

}

func (h *Heap) Insert(n *Node) {
	h.nodes = append(h.nodes, n)
	h.HeapifyUp(len(h.nodes) - 1)

}

func (h *Heap) Size() int {
	return len(h.nodes)
}
