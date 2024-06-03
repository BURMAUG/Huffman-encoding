package heap

type Heap struct {
	nodes []*Node
}

// Parent returns the parent node index
func Parent(index int) int {
	return (index - 1) / 2
}

// LeftChild returns the left child index of a given node.
func LeftChild(index int) int {
	return 2*index + 1
}

// RightChild returns the right child index of a given node.
func RightChild(index int) int {
	return 2*index + 2
}

// HasLeftChild checks to see if the node has a left child or not.
func (h *Heap) HasLeftChild(index int) bool {
	return LeftChild(index) < len(h.nodes)
}

// HasRightChild checks to see if the current node has a right child.
func (h *Heap) HasRightChild(index int) bool {
	return RightChild(index) < len(h.nodes)
}

// Swap interchange two nodes in a heap.
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
