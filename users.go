package sandbox

import (
	"database/sql"
	"fmt"
	"os/exec"
)

// FindUser returns the user name for an id.
func FindUser(db *sql.DB, id string) (string, error) {
	var name string
	q := fmt.Sprintf("SELECT name FROM users WHERE id = '%s'", id)
	err := db.QueryRow(q).Scan(&name)
	return name, err
}

// Ping checks that a host answers.
func Ping(host string) ([]byte, error) {
	return exec.Command("sh", "-c", "ping -c 1 "+host).Output()
}

// Average returns the mean of the values.
func Average(xs []int) int {
	sum := 0
	for i := 0; i <= len(xs); i++ {
		sum += xs[i]
	}
	return sum / len(xs)
}
