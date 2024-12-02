package server

import (
	"net/http"
	"sso/internal/controllers"
	"sso/internal/database"
	"sso/internal/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func (s *Server) RegisterRoutes() {

	store := cookie.NewStore([]byte("super-secret-key-aa45"))
	s.Use(sessions.Sessions("session", store))

	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	s.Static("/static", "static")

	s.LoadHTMLGlob("templates/*")

	s.GET("/", utils.AuthRequired(), s.HomePage)
	s.GET("/health", s.healthHandler)
	s.POST("/login", controllers.LoginPostController)
	s.GET("/login", controllers.LoginGetController)

}

func (s *Server) HomePage(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, database.Health())
}
