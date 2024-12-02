package database

import (
	"context"
	"database/sql"
	"errors"
	"sso/internal/models"
	"strings"
)

func GetClient(ctx context.Context, clientID string) (models.Client, error) {
	var client models.Client
	var redirectURIs string
	client.ClientID = clientID

	query := `SELECT id, email, name, allowed_redirect_uris FROM clients WHERE id = $1`
	err := db.QueryRowContext(ctx, query, clientID).Scan(&client.ID, &client.Email, &client.Name, &redirectURIs)
	if err != nil {
		if err == sql.ErrNoRows {
			return client, errors.New("client not found")
		}
		return client, err
	}

	client.AllowedRedirectURIs = parseRedirectURIs(redirectURIs)
	return client, nil
}

func parseRedirectURIs(redirectURIs string) []string {
	if redirectURIs == "" {
		return []string{}
	}
	return strings.Split(redirectURIs, ",")
}
