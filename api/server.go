package api

import (
	"fmt"
	// "go/token"
	db "github.com/21toffy/relinc/db/sqlc"

	"github.com/21toffy/relinc/token"
	"github.com/21toffy/relinc/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

// serve all HTTP request for banking service// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      *store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil

}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user-and-account", server.CreateUserAccount)
	router.POST("/user/login", server.LoginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts-for-user", server.CreateAccount)

	authRoutes.GET("/accounts/:id", server.GetAccount)
	authRoutes.GET("/all/accounts", server.ListAccount)
	authRoutes.POST("/transfers", server.TransferMoney)
	server.router = router

}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	fmt.Println(err)
	return gin.H{"error": err.Error()}
}

func DbErrorResponse(err *pq.Error) gin.H {
	switch err.Code.Name() {
	case "unique_violation":
		return gin.H{"error": "A unique constraint violation has occurred. Please check that the values you are entering are unique."}
	case "foreign_key_violation":
		return gin.H{"error": "A foreign key violation has occurred. Please check that the values you are entering are valid."}
	case "not_null_violation":
		return gin.H{"error": "A non-null constraint violation has occurred. Please check that all required fields are filled in."}
	case "check_violation":
		return gin.H{"error": "A check constraint violation has occurred. Please check that the values you are entering are valid."}
	case "deadlock_detected":
		return gin.H{"error": "A deadlock has been detected. Please try your operation again later."}
	case "serialization_failure":
		return gin.H{"error": "A serialization failure has occurred. Please try your operation again later."}
	case "syntax_error":
		return gin.H{"error": "A syntax error has been detected in your SQL statement."}
	case "undefined_function":
		return gin.H{"error": "A function in your SQL statement is not defined in the database."}
	case "invalid_cursor_state":
		return gin.H{"error": "An invalid cursor state has been detected."}
	case "invalid_transaction_state":
		return gin.H{"error": "An invalid transaction state has been detected."}
	default:
		return gin.H{"error": "An unknown error has occurred."}
	}
}
