package api

import (
	"fmt"
	db "relinc/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

// serve all HTTP request for banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrrency)
	}

	router.POST("/user-and-account", server.CreateUserAccount)
	router.POST("/accounts-for-user", server.CreateAccount)

	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.ListAccount)
	router.POST("/transfers", server.TransferMoney)

	server.router = router
	return server
}

// start the HTTP server
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
