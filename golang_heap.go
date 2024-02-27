package dsa

type LessFunc[T comparable] func(val1, val2 T) bool

type PriorityQueue[T comparable] struct {
	data []T
	less LessFunc[T]
}

func NewPriorityQueue[T comparable](less LessFunc[T]) *PriorityQueue[T] {
	if less == nil {
		panic("Less function is nil")
	}
	return &PriorityQueue[T]{
		data: make([]T, 0),
		less: less,
	}
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq.data)
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.data[i], pq.data[j])
}

func (pq *PriorityQueue[T]) Push(val any) {
	pq.data = append(pq.data, val.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	val := pq.data[pq.Len()-1]
	pq.data = pq.data[:pq.Len()-1]
	return val
}

func (pq *PriorityQueue[T]) Peek() any {
	return pq.data[0]
}
