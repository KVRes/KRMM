package KRU

type Stack struct {
	layers []*Layer
}

func (s *Stack) Push(layer *Layer) {
	s.layers = append(s.layers, layer)
}

func (s *Stack) PushFrame() {
	s.Push(NewLayer(LayerType_StackFrame))
}

func (s *Stack) Top() *Layer {
	if len(s.layers) == 0 {
		return nil
	}
	return s.layers[len(s.layers)-1]
}

func (s *Stack) Pop() *Layer {
	if len(s.layers) == 0 {
		return nil
	}
	layer := s.layers[len(s.layers)-1]
	s.layers = s.layers[:len(s.layers)-1]
	return layer
}

func (s *Stack) Raw() []*Layer {
	return s.layers
}

func (s *Stack) Len() int {
	return len(s.layers)
}
