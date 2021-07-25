package api

import (
	"fmt"

	"github.com/shawon1fb/go_api/middle_ware"

	db "github.com/shawon1fb/go_api/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shawon1fb/go_api/token"
	"github.com/shawon1fb/go_api/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	filter     *util.CuckooFilter
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		filter:     util.NewCuckooFilter(),
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// router.POST("/users", server.createUser)
	router.POST("/", server.helloWorld)
	router.POST("/users", server.createUser)
	router.POST("/users/login", middle_ware.LoginUserFilterMiddleware(server.filter), server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/getUser", server.getUser)

	// authRoutes.POST("/accounts", server.createAccount)
	// authRoutes.GET("/accounts/:id", server.getAccount)
	// authRoutes.GET("/accounts", server.listAccounts)

	// authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
