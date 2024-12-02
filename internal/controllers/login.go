package controllers

import (
	"context"
	"net/http"
	"sso/internal/database"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginGetController(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginPostController(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hashedPassword, err := database.GetPassword(ctx, email)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
		return
	}

	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Set("email", email)
	err = session.Save()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Could not save session"})
		return
	}

}
