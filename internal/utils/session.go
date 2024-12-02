package utils

import (
	"fmt"
	"net/http"
	"sso/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VerifySession(c *gin.Context) (models.UserModel, error) {
	session := sessions.Default(c)
	authenticated, authOK := session.Get("authenticated").(bool)
	email, emailOK := session.Get("email").(string)
	name, nameOK := session.Get("name").(string)
	isAdmin, adminOK := session.Get("isAdmin").(string)

	if !adminOK {
		isAdmin = "false"
	}

	if !authOK || !authenticated || !emailOK || !nameOK {
		return models.UserModel{
			Name:    name,
			Email:   email,
			IsAdmin: isAdmin == "true",
		}, fmt.Errorf("unauthorized user")
	}

	return models.UserModel{
		Name:    name,
		Email:   email,
		IsAdmin: isAdmin == "true",
	}, nil

}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authenticated, ok := session.Get("authenticated").(bool)
		if !ok || !authenticated {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
