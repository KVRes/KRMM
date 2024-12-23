package KRU

type LayerType int

const (
	LayerType_Heap LayerType = iota - 1
	LayerType_Stack
	LayerType_StackFrame
)

type Layer struct {
	m map[string]*Element
	t LayerType
}

func NewLayer(t LayerType) *Layer {
	return &Layer{
		m: make(map[string]*Element),
		t: t,
	}
}

func (l *Layer) Add(key string, data any) {
	l.m[key] = &Element{Data: data}
}

func (l *Layer) Get(key string) *Element {
	return l.m[key]
}

func (l *Layer) Remove(key string) {
	delete(l.m, key)
}

type Element struct {
	Data any
}

func (l *Layer) Type() LayerType {
	return l.t
}
