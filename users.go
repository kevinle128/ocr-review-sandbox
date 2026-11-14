package sandbox

import (
	"database/sql"
	"errors"
	"os/exec"
)

// FindUser returns the user name for an id.
func FindUser(db *sql.DB, id string) (string, error) {
	var name string
	err := db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)
	return name, err
}

// Ping checks that a host answers.
func Ping(host string) ([]byte, error) {
	return exec.Command("ping", "-c", "1", "--", host).Output()
}

// Average returns the mean of the values.
func Average(xs []int) (int, error) {
	if len(xs) == 0 {
		return 0, errors.New("no values")
	}
	sum := 0
	for _, x := range xs {
		sum += x
	}
	return sum / len(xs), nil
}

// Max returns the largest value, or 0 for an empty slice.
func Max(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	m := xs[0]
	for _, x := range xs[1:] {
		if x > m {
			m = x
		}
	}
	return m
}
