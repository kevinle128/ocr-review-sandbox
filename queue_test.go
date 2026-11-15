package sandbox

import (
	"sync"
	"testing"
)

func TestQueueFIFO(t *testing.T) {
	var q Queue
	if got, ok := q.Pop(); ok || got != "" {
		t.Fatalf("Pop on empty = (%q, %v), want (\"\", false)", got, ok)
	}
	for _, s := range []string{"a", "b", "c"} {
		q.Push(s)
	}
	if got := q.Len(); got != 3 {
		t.Fatalf("Len = %d, want 3", got)
	}
	for _, want := range []string{"a", "b", "c"} {
		got, ok := q.Pop()
		if !ok || got != want {
			t.Fatalf("Pop = (%q, %v), want (%q, true)", got, ok, want)
		}
	}
	if got, ok := q.Pop(); ok || got != "" {
		t.Fatalf("Pop after drain = (%q, %v), want (\"\", false)", got, ok)
	}
	if got := q.Len(); got != 0 {
		t.Fatalf("Len = %d, want 0", got)
	}
}

func TestQueueConcurrent(t *testing.T) {
	const workers, perWorker = 8, 200
	var q Queue
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < perWorker; j++ {
				q.Push("x")
				_ = q.Len()
			}
		}()
	}
	wg.Wait()
	if got := q.Len(); got != workers*perWorker {
		t.Fatalf("Len = %d, want %d", got, workers*perWorker)
	}

	var mu sync.Mutex
	popped := 0
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if _, ok := q.Pop(); !ok {
					return
				}
				mu.Lock()
				popped++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	if popped != workers*perWorker {
		t.Fatalf("popped %d, want %d", popped, workers*perWorker)
	}
}
