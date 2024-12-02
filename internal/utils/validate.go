package utils

import (
	"context"
	"sso/internal/database"
	"sso/internal/models"
)



// if return 0 then invalid client id
// if 1 then invalid redirect uri
// if 2 then its validated
func SSOValidator(clientID string, redirectURI string, ctx context.Context) int {
	client, err := database.GetClient(ctx, clientID)

	if err != nil || client.ID == "" {
		return 0
	}

	if !isRedirectURIPermitted(client, redirectURI) {
		return 1
	}

	return 2
}

func isRedirectURIPermitted(client models.Client, redirectURI string) bool {
	for _, uri := range client.AllowedRedirectURIs {
		if uri == redirectURI {
			return true
		}
	}
	return false
}
