package controller

import (
	"net/http"

	"github.com/DeveloperGerald/TurtleSoup/service"
	"github.com/gin-gonic/gin"
)

type GiveAnswerReq struct {
	StoryID    string `json:"story_id" required:"true"`
	UserAnswer string `json:"user_answer" required:"true"`
}

type GiveAnswerRespData struct {
	Answer string `json:"answer"`
}

func GiveAnswerController(c *gin.Context) {
	var req GiveAnswerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	answer, err := service.GiveAnswer(req.StoryID, req.UserAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StandardResp{Code: 0, Data: &GiveAnswerRespData{Answer: answer}})
}

func GetRandomStoryController(c *gin.Context) {
	story, err := service.GetRandomStory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, StandardResp{Code: -1, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StandardResp{Code: 0, Data: story})
}
