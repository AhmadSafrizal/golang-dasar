package handler

import (
	"log"
	"net/http"

	"github.com/AhmadSafrizal/golang-dasar/tugas7/helper"
	model "github.com/AhmadSafrizal/golang-dasar/tugas7/models"
	repository "github.com/AhmadSafrizal/golang-dasar/tugas7/repositories"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repository *repository.UserRepository
}

func (u *UserHandler) GetGorm(ctx *gin.Context) {
	users, err := u.Repository.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserHandler) CreateGorm(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}

	if !helper.IsValidEmail(user.Email) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid email format",
		})
		return
	}

	if len(user.Password) < 6 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "password must be at least 6 characters",
		})
		return
	}

	userFetched, err := u.Repository.GetByEmail(user.Email)
	if err == nil && userFetched.ID != 0 {
		ctx.AbortWithStatusJSON(http.StatusConflict, map[string]any{
			"message": "email already registered",
		})
		return
	}

	passwordHashed, err := helper.HashPassword(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	user.Password = passwordHashed

	err = u.Repository.Create(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	// bind body
	user := &model.User{}
	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]any{
			"message": "invalid body request",
		})
		return
	}

	userFetched, err := u.Repository.GetByEmail(user.Email)
	if err != nil || userFetched.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
			"message": "email not found",
		})
		return
	}
	valid := helper.CheckPasswordHash(user.Password, userFetched.Password)
	if !valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]any{
			"message": "wrong password",
		})
		return
	}
	token, err := helper.GenerateUserJWT(userFetched.Username, userFetched.Email)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"message": "something went wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}
