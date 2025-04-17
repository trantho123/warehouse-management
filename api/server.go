package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/trantho123/warehouse-management/db/sqlc"
	"github.com/trantho123/warehouse-management/utils"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config utils.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
