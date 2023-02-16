package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/21toffy/relinc/db/sqlc"
	"github.com/21toffy/relinc/token"
	"github.com/21toffy/relinc/util"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// example_dob = "2023-02-04T09:35:25.490276Z"
type CreateUserAccountRequest struct {
	Password       string    `json:"password" binding:"required,min=6"`
	EmailAddress   string    `json:"email_address" binding:"required"`
	FirstName      string    `json:"first_name" binding:"required"`
	LastName       string    `json:"last_name" binding:"required"`
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

type CreateUserAccountRsponse struct {
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

func (server *Server) CreateUserAccount(ctx *gin.Context) {
	var req CreateUserAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("the error ocured here")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserAccountTxParams{
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
		Password:       hashedPassword,
	}

	account, err := server.store.CreateUserAccountTx(ctx, arg)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			ctx.JSON(http.StatusBadRequest, DbErrorResponse(pqError))

		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := CreateUserAccountRsponse{
		FirstName:      account.User.FirstName,
		LastName:       account.User.LastName,
		EmailAddress:   account.User.EmailAddress,
		PhoneNumber:    account.User.PhoneNumber,
		Username:       account.User.Username,
		Dob:            account.User.Dob,
		Address:        account.User.Address,
		ProfilePicture: account.User.ProfilePicture,
		Gender:         account.User.Gender,
		Balance:        account.Account.Balance,
		Currency:       account.Account.Currency,
		AccountType:    account.Account.AccountType,
	}
	ctx.JSON(http.StatusOK, res)
	return
}

type PublicUserResponse struct {
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	EmailAddress string `json:"email_address" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
	Username     string `json:"username" binding:"required"`
}

func newUserResponse(user db.User) PublicUserResponse {
	return PublicUserResponse{
		Username:     user.Username,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		EmailAddress: user.EmailAddress,
		PhoneNumber:  user.PhoneNumber,
	}
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

	account, err := server.store.GetUserAccountWithUserFields(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Email != account.EmailAddress {
		err = fmt.Errorf("Account does not belong to owner")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
	return
}

func (server *Server) ListAccount(ctx *gin.Context) {

	// if err := ctx.ShouldBindUri(&req); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	// 	return
	// }
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	accounts, err := server.store.GetUsersAccountsByUserEmail(ctx, authPayload.Email)

	// arg := db.GetUsersAccountsParams{
	// 	Email: user.,
	// 	Limit:  req.PageSize,
	// 	Offset: (req.PageID - 1) * req.PageSize,
	// }

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return

	}
	ctx.JSON(http.StatusOK, accounts)
	return
}

type CreateAccountRequest struct {
	// Owner       int64  `json:"owner" binding:"required"`
	Balance     int64  `json:"balance" binding:"required"`
	Currency    string `json:"currency" binding:"required,currency"`
	AccountType string `json:"account_type" binding:"required"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("the error ocured here")
		fmt.Println(req, 1234)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUserByEmail(ctx, authPayload.Email)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserAccountParams{
		Owner:       user.ID,
		Balance:     req.Balance,
		Currency:    req.Currency,
		AccountType: req.AccountType,
	}

	account, err := server.store.CreateUserAccount(ctx, arg)
	if err != nil {
		if pqError, ok := err.(*pq.Error); ok {
			ctx.JSON(http.StatusBadRequest, DbErrorResponse(pqError))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
	return
}

type LoginUserRequest struct {
	Password     string `json:"password" binding:"required,min=6"`
	EmailAddress string `json:"email_address" binding:"required"`
}

type LoginUserRsponse struct {
	SessionID             uuid.UUID          `json:"session_id"`
	AccessToken           string             `json:"access_token"`
	AccessTokenExpiresAt  time.Time          `json:"access_token_expires_at"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time          `json:"refresh_token_expires_at"`
	User                  PublicUserResponse `json:"user"`
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.EmailAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.EmailAddress,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.EmailAddress,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := LoginUserRsponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}
