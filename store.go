package sandbox

import "errors"

// Store keeps account balances in memory.
type Store struct {
	balances map[string]int
}

// NewStore returns an empty store.
func NewStore() *Store {
	return &Store{balances: map[string]int{}}
}

// Balance returns the balance of an account.
func (s *Store) Balance(id string) (int, error) {
	b, ok := s.balances[id]
	if !ok {
		return 0, errors.New("unknown account")
	}
	return b, nil
}
