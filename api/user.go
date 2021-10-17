package api

import (
	//"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "github.com/daryalazareva/microapp/db/sqlc"
	"github.com/daryalazareva/microapp/util"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//handler for signup
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	encryptedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:             req.Email,
		EncryptedPassword: encryptedPassword,
	}

	var u db.User
	u, err = server.store.CreateRecordUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "ID: "+strconv.Itoa(int(u.ID)))
}

//handler for signin
func (server *Server) authUser(ctx *gin.Context) {

	email, password, ok := ctx.Request.BasicAuth()
	if email == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("email should not be empty")))
		return
	}
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("basic auth failed")))
		return
	}

	u, err := server.store.GetRecordUser(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Incorrect email or password")))
		return
	}

	err = util.CheckPassword(password, u.EncryptedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("Incorrect email or password")))
		return
	}

	token, err := server.GenerateToken(email, u.EncryptedPassword)

	ctx.JSON(http.StatusOK, token+"	"+time.Now().String())
}

type UpdateUserRequest struct {
	NewPassword string `json:"newpassword" binding:"required"`
}

//handler for changepassword
func (server *Server) updateUsersPassword(ctx *gin.Context) {

	bearerToken := ctx.Request.Header.Get("Authorization")
	var authorizationToken string
	tokenSlice := strings.Split(bearerToken, " ")

	if len(tokenSlice) < 2 {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid token")))
		return
	}
	if tokenSlice[0] != "Bearer" {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Not JWT")))
		return
	}
	authorizationToken = tokenSlice[1]
	email, err := server.VerifyToken(authorizationToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(errors.New("invalid token")))
		return
	}

	var req UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	encryptedPassword, err := util.HashPassword(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.ChangePasswordTx(ctx, email, encryptedPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "ok")
}
