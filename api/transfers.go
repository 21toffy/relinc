package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "relinc/db/sqlc"

	"github.com/gin-gonic/gin"
)

type TransferMoneyRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required"`
	ToAccountID   int64  `json:"to_account_id" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) TransferMoney(ctx *gin.Context) {
	var req TransferMoneyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fromAccountValidity, fromAccountBalance := server.validAccount(ctx, req.FromAccountID, req.Currency, "from")
	if !fromAccountValidity {
		return
	}
	toAccountValidity, _ := server.validAccount(ctx, req.ToAccountID, req.Currency, "to")
	if !toAccountValidity {
		return
	}

	if req.Amount > fromAccountBalance {
		err := fmt.Errorf("insufficient funds")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.FromAccountID,
		ToAccountID:   req.ToAccountID,
		Amount:        req.Amount,
	}

	account, err := server.store.TransferTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
	return
}

func (server *Server) validAccount(ctx *gin.Context, accountId int64, currency, acctLabel string) (bool, int64) {
	account, err := server.store.GetUsertAccountByAccountId(ctx, accountId)
	if err != nil {
		if err == sql.ErrNoRows {
			err := fmt.Errorf("%s account not found", acctLabel)
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false, 0
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false, 0

	}
	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch from: %v account to %s account", account.Owner, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false, 0
	}
	return true, account.Balance
}
