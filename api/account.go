package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "relinc/db/sqlc"
	"time"

	"github.com/gin-gonic/gin"
)

// example_dob = "2023-02-04T09:35:25.490276Z"
type CreateAccountRequest struct {
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
	EmailAddress   string    `json:"email_address" binding:"required"`
	PhoneNumber    string    `json:"phone_number" binding:"required"`
	Username       string    `json:"username" binding:"required"`
	Dob            time.Time `json:"dob" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	Address        string    `json:"address" binding:"required"`
	ProfilePicture string    `json:"profile_picture" binding:"required"`
	Gender         string    `json:"gender" binding:"required"`
	Balance        int64     `json:"balance" binding:"required"`
	Currency       string    `json:"currency" binding:"required,currency"`
	AccountType    string    `json:"account_type" binding:"required"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("the error ocured here")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountTxParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		EmailAddress:   req.EmailAddress,
		PhoneNumber:    req.PhoneNumber,
		Username:       req.Username,
		Dob:            req.Dob,
		Address:        req.Address,
		ProfilePicture: req.ProfilePicture,
		Gender:         req.Gender,
		Balance:        req.Balance,
		Currency:       req.Currency,
		AccountType:    req.AccountType,
	}

	account, err := server.store.CreateAccountTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
	return
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccountTx(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, account)
	return
}

type ListAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) ListAccount(ctx *gin.Context) {

	var req ListAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetUsersAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.GetUsersAccounts(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, accounts)
	return
}
