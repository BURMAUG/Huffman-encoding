package heap

type HeapInterface interface {
	Peek() *Node
	DeleteMin() *Node
	Insert(n Node)
	Size() int32
}

type Heap struct {
	nodes []*Node
}
