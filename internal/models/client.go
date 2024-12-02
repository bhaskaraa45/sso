package models

type Client struct {
	ID                  string   `json:"id"`
	ClientID            string   `json:"client_id"`
	Name                string   `json:"name"`
	Email               string   `json:"email"`
	AllowedRedirectURIs []string `json:"allowed_redirect_uris"`
}
