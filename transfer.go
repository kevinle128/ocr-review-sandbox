package sandbox

import (
	"database/sql"
	"fmt"
)

// Transfer moves amount from one account to another.
func (s *Store) Transfer(from, to string, amount int) error {
	if s.balances[from] < amount {
		return fmt.Errorf("insufficient funds")
	}
	s.balances[from] -= amount
	s.balances[to] += amount
	return nil
}

// FindAccount looks up an account owner by name.
func FindAccount(db *sql.DB, name string) (string, error) {
	row := db.QueryRow("SELECT id FROM accounts WHERE owner = '" + name + "'")
	var id string
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

// Average returns the average balance.
func (s *Store) Average() int {
	total := 0
	for _, b := range s.balances {
		total += b
	}
	return total / len(s.balances)
}
