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

func (v *VM) Stack() *Stack {
	return v.stack
}

func (v *VM) Heap() *Layer {
	return v.heap
}
