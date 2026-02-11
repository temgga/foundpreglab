package heap

type CompareResult int8

const (
	Less    CompareResult = -1
	Greater CompareResult = 1
	Eq      CompareResult = 0
)

type Heap struct {
	data       []int
	comparator func(a, b int) CompareResult
}

func NewHeap(comparator func(int, int) CompareResult) *Heap {
	return &Heap{
		data:       make([]int, 0),
		comparator: comparator,
	}
}

func (h *Heap) Push(value int) {
	h.data = append(h.data, value)
	h.bubbleUp(len(h.data) - 1)
}

func (h *Heap) bubbleUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.comparator(h.data[index], h.data[parent]) >= 0 {
			break
		}
		h.data[index], h.data[parent] = h.data[parent], h.data[index]
		index = parent
	}
}

func (h *Heap) Len() int {
	return len(h.data)
}

func (h *Heap) IsEmpty() bool {
	return len(h.data) == 0
}
