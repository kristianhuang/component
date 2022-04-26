package slicequeue

import "sync"

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(cap int) *SliceQueue {
	return &SliceQueue{data: make([]interface{}, 0, cap)}
}

// Push is inbound queue.
func (q *SliceQueue) Push(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Pop is out of queue.
func (q *SliceQueue) Pop() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()

	return v
}
