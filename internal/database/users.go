package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sso/internal/models"
)

func GetPassword(ctx context.Context, email string) (string, error) {
	var passwordHash string

	query := "SELECT password_hash FROM users WHERE email = $1"

	err := db.QueryRowContext(ctx, query, email).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no user found with email: %s", email)
		}
		return "", fmt.Errorf("error querying password for email %s: %v", email, err)
	}

	return passwordHash, nil
}

func GetUser(ctx context.Context, email string) (models.UserModel, error) {
	var user models.UserModel
	query := "SELECT name, email, is_admin FROM users WHERE email = $1"

	err := db.QueryRowContext(ctx, query, email).Scan(&user.Name, &user.Email, &user.IsAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user not found")
		}
		return user, err
	}

	return user, nil
}
