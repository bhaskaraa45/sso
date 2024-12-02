package utils

import (
	"context"
	"fmt"
	"net/url"
	"sso/internal/database"
	"sso/internal/models"
)

// if return 0 then invalid client id
// if 1 then invalid redirect uri
// if 2 then its validated
func SSOValidator(clientID string, redirectURI string, ctx context.Context) int {
	client, err := database.GetClient(ctx, clientID)

	fmt.Println(client.AllowedRedirectURIs)
	fmt.Println(redirectURI)

	if err != nil || client.ID == "" {
		return 0
	}

	if !isRedirectURIPermitted(client, redirectURI) {
		return 1
	}

	return 2
}

func isRedirectURIPermitted(client models.Client, redirectURI string) bool {
	providedURI, err := url.Parse(redirectURI)
	if err != nil {
		return false
	}

	for _, uri := range client.AllowedRedirectURIs {
		storedURI, err := url.Parse(uri)

		fmt.Println(err)

		
		if err == nil && storedURI.String() == providedURI.String() {
			return true
		}
	}
	return false
}
