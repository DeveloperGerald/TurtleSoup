package controller

import (
	"net/http"

	"github.com/DeveloperGerald/TurtleSoup/consts"
	myjwt "github.com/DeveloperGerald/TurtleSoup/pkg/jwt"
	"github.com/DeveloperGerald/TurtleSoup/service"
	"github.com/gin-gonic/gin"
)

type CreateStoryReq struct {
	Title     string `json:"title"`
	Riddle    string `json:"riddle"`
	FullStory string `json:"full_story"`
}

func CreateStoryController(c *gin.Context) {
	var req CreateStoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	userInfo, exists := c.Get(consts.UserClaimsKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, StandardResp{Code: -1, Message: "Got no user info"})
		return
	}

	userClaims, ok := userInfo.(*myjwt.UserClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, StandardResp{Code: -1, Message: "Got no user info"})
		return
	}

	story, err := service.CreateStory(req.Title, req.Riddle, req.FullStory, userClaims.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, StandardResp{Code: 0, Data: story})
}
