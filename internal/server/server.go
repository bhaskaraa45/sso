package server

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"sso/internal/database"
)

type Server struct {
	port int
	db   *sql.DB
	*gin.Engine
}

func NewServer() *Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	server := &Server{
		port:   port,
		db:     database.InitializeDB(),
		Engine: gin.Default(),
	}

	server.RegisterRoutes()

	return server
}
