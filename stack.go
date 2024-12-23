package KRU

import (
	"sync"
)

type Stack struct {
	l      sync.RWMutex
	prev   *Stack
	layers []*Layer
}

func (s *Stack) QueryScope(key string) *Element {
	return s.query(key, true)
}

func (s *Stack) QueryGlobal(key string) *Element {
	return s.query(key, false)
}

func (s *Stack) query(key string, stopAtStackFrame bool) *Element {
	if s == nil {
		return nil
	}
	s.l.RLock()
	defer s.l.RUnlock()
	for i := len(s.layers) - 1; i >= 0; i-- {
		layer := s.layers[i]
		switch layer.Type() {
		case LayerType_StackFrame:
			if stopAtStackFrame {
				return nil
			}
		default:
			elem, _ := layer.m.Get(key)
			if elem != nil {
				return elem
			}
		}
	}
	if s.prev != nil {
		return s.prev.query(key, stopAtStackFrame)
	}
	return nil
}

func (s *Stack) Fork() *Stack {
	return &Stack{
		prev: s,
	}
}

func (s *Stack) Push(layer *Layer) {
	s.l.Lock()
	defer s.l.Unlock()
	s.layers = append(s.layers, layer)
}

func (s *Stack) PushFrame() {
	s.Push(NewLayer(LayerType_StackFrame))
}

func (s *Stack) Top() *Layer {
	s.l.RLock()
	defer s.l.RUnlock()
	t := s._top()
	if t != nil {
		return t
	}
	if s.prev != nil {
		return s.prev.Top()
	}
	return nil
}

func (s *Stack) SelfTop() *Layer {
	s.l.RLock()
	defer s.l.RUnlock()
	return s._top()
}

func (s *Stack) _top() *Layer {
	if len(s.layers) == 0 {
		return nil
	}
	return s.layers[len(s.layers)-1]
}

func (s *Stack) Pop() *Layer {
	s.l.Lock()
	defer s.l.Unlock()
	if len(s.layers) == 0 {
		return nil
	}
	layer := s.layers[len(s.layers)-1]
	s._pop()
	return layer
}

func (s *Stack) _pop() {
	if len(s.layers) == 0 {
		return
	}
	s.layers = s.layers[:len(s.layers)-1]
}

func (s *Stack) PopFrame() {
	s.l.Lock()
	defer s.l.Unlock()
	if len(s.layers) == 0 {
		return
	}
	for len(s.layers) > 0 && s._top().Type() != LayerType_StackFrame {
		s._pop()
	}
	if len(s.layers) == 0 {
		return
	}
	s._pop()
}

func (s *Stack) Raw() []*Layer {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.layers
}

func (s *Stack) Len() int {
	s.l.RLock()
	defer s.l.RUnlock()
	return len(s.layers)
}
