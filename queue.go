package sandbox

import "sync"

// Queue is a FIFO queue safe for concurrent use.
type Queue struct {
	mu    sync.Mutex
	items []string
}

// Push adds an item to the back.
func (q *Queue) Push(s string) {
	q.items = append(q.items, s)
}

// Pop removes and returns the front item.
func (q *Queue) Pop() string {
	q.mu.Lock()
	s := q.items[0]
	q.items = q.items[1:]
	q.mu.Unlock()
	return s
}

// Len returns the number of items.
func (q *Queue) Len() int {
	return len(q.items)
}
