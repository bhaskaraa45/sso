package controllers

import (
	"context"
	"fmt"
	"net/http"
	"sso/internal/database"
	"sso/internal/utils"
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

	redirectURI := c.Query("redirect_uri")

	if redirectURI == "" {
		if user.IsAdmin {
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"AccessRestricted": true,
			})
		}
		return
	}

	clientID := c.Query("client_id")

	if clientID == "" {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error":            "Missing Client ID.",
			"DeveloperMessage": "Developers: Please include the query parameter client_id in the login URL. Example: https://sso.bhaskaraa45.me/login?client_id=your_client_id",
		})
		return
	}



	res := utils.SSOValidator(clientID, redirectURI, ctx)

	if res == 0 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error":            "Invalid Client ID.",
			"DeveloperMessage": "Developers: The client_id provided in the query parameter is not valid. Verify the client_id or check your client configuration.",
		})
		return
	}

	if res == 1 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error":            "Unregistered Redirect URI.",
			"DeveloperMessage": "Developers: The redirect_uri provided in the query parameter is not registered for this client ID. Please ensure the URI is added to the client configuration.",
		})
		return
	}

	if res != 2 {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"Error": "Unexpected Error",
		})
		return
	}

}
