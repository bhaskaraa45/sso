package controllers

import (
	"context"
	"fmt"
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

	// Fetch user details from the database
	user, err := database.GetUser(c.Request.Context(), email)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "User not found"})
		return
	}

	hashedPassword, err := database.GetPassword(ctx, email)
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"Error": "Invalid credentials"})
		return
	}

	// If credentials are correct, start a session
	session := sessions.Default(c)
	session.Set("authenticated", true)
	session.Set("email", user.Email)
	session.Set("name", user.Name)
	session.Set("is_admin", user.IsAdmin)
	err = session.Save()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{"Error": "Could not save session"})
		return
	}

	if user.IsAdmin {
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"AccessRestricted": true,
		})
	}
}
