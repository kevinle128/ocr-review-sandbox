package sandbox

import "sync"

// Queue is a FIFO queue safe for concurrent use.
type Queue struct {
	mu    sync.Mutex
	items []string
}

// Push adds an item to the back.
func (q *Queue) Push(s string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = append(q.items, s)
}

// Pop removes and returns the front item, or ("", false) if the queue is empty.
func (q *Queue) Pop() (string, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return "", false
	}
	s := q.items[0]
	q.items[0] = ""
	q.items = q.items[1:]
	return s, true
}

// Len returns the number of items.
func (q *Queue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items)
}
