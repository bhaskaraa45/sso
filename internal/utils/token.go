package utils

import (
	"os"
	"sso/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("TOKEN_SECRET"))

// GenerateToken generates a JWT
func GenerateToken(issuer, audience, clientID, hostedDomain string, user models.UserModel) (string, error) {
	claims := jwt.MapClaims{
		"iss":            issuer,                               // Token issuer
		"aud":            audience,                             // Intended audience (e.g., client app)
		"sub":            user.Id,                              // Subject (unique user identifier)
		"email":          user.Email,                           // User's email address
		"email_verified": true,                                 // Whether the email is verified
		"hd":             hostedDomain,                         // Hosted domain (optional)
		"name":           user.Name,                            // User's full name
		"iat":            time.Now().Unix(),                    // Issued at (current time)
		"exp":            time.Now().Add(1 * time.Hour).Unix(), // Expiration time (1 hour)
		"client_id":      clientID,                             // Client ID
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
