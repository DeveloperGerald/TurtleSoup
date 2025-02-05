package controller

import (
	"net/http"

	"github.com/DeveloperGerald/TurtleSoup/service"
	"github.com/gin-gonic/gin"
)

type RegisterUserReq struct {
	Name     string  `json:"name" binding:"required"`
	Password string  `json:"password" binding:"required,min=8"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

type LoginUserReq struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

func RegisterUserController(c *gin.Context) {
	var req RegisterUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	user, err := service.RegisterUser(req.Name, req.Password, req.Email, req.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, StandardResp{Code: 0, Data: user})
}

func LoginUserController(c *gin.Context) {
	var req LoginUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	token, err := service.LoginUser(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StandardResp{Code: 0, Data: gin.H{"token": token}})
}
