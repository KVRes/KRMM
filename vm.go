package KRU

type VM struct {
	stack *Stack
	heap  *Layer
}

func NewVM() *VM {
	return &VM{
		stack: &Stack{},
		heap:  NewLayer(LayerType_Heap),
	}
}
