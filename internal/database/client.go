package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sso/internal/models"

	"github.com/lib/pq"
)

func GetClient(ctx context.Context, clientID string) (models.Client, error) {
	var client models.Client
	client.ClientID = clientID
	var redirectURIs []string

	tx, err := db.Begin()
	if err != nil {
		return client, err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `SELECT id, email, name, redirected_uris FROM clients WHERE client_id = $1`
	err = tx.QueryRow(query, clientID).Scan(&client.ID, &client.Email, &client.Name, pq.Array(&redirectURIs))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return client, fmt.Errorf("client not found for client_id: %s", clientID)
		}
		return client, fmt.Errorf("error fetching client: %w", err)
	}

	client.AllowedRedirectURIs = redirectURIs
	return client, nil
}
