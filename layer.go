package KRMM

import omap "github.com/wk8/go-ordered-map/v2"

type LayerType int

const (
	LayerType_Heap LayerType = iota - 1
	LayerType_Stack
	LayerType_StackFrame
)

type Layer struct {
	m *omap.OrderedMap[string, *Element]
	t LayerType
}

func NewLayer(t LayerType) *Layer {
	return &Layer{
		m: omap.New[string, *Element](),
		t: t,
	}
}

func (l *Layer) Set(key string, data any) {
	l.m.Set(key, &Element{Data: data})
}

func (l *Layer) Get(key string) *Element {
	elem, _ := l.m.Get(key)
	return elem
}

func (l *Layer) Remove(key string) {
	l.m.Delete(key)
}

func (l *Layer) Len() int {
	return l.m.Len()
}

func (l *Layer) Raw() *omap.OrderedMap[string, *Element] {
	return l.m
}

type Element struct {
	Data any
}

func (l *Layer) Type() LayerType {
	return l.t
}
