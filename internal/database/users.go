package database

import (
	"context"
	"database/sql"
	"fmt"
)

func GetPassword(ctx context.Context, email string) (string, error) {
	var passwordHash string

	query := "SELECT password FROM users WHERE email = $1"

	err := db.QueryRowContext(ctx, query, email).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with email: %s", email)
		}
		return "", fmt.Errorf("error querying password for email %s: %v", email, err)
	}

	return passwordHash, nil
}
