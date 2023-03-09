package quiz

import (
	"github.com/gin-gonic/gin"
	"quiz/model"
	response "quiz/model"
	"quiz/service"
)

type CommentApi struct {
}

var commentService = service.ServiceGroupApp.CommentService

func (commentApi *CommentApi) CreateComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.CreateComment(comment); err != nil {
		response.FailWithMessage("新增失敗", c)
	}

	response.OkWithMessage("新增成功", c)
}

func (commentApi *CommentApi) DeleteComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.DeleteComment(comment); err != nil {
		response.FailWithMessage("刪除失敗", c)
	} else {
		response.OkWithMessage("刪除成功", c)
	}
}

func (commentApi *CommentApi) UpdateComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := commentService.UpdateComment(comment); err != nil {
		response.FailWithMessage("更新失敗", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindComment 用UUID查詢
func (commentApi *CommentApi) FindComment(c *gin.Context) {
	var comment model.Comment
	err := c.ShouldBindQuery(&comment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recomment, err := commentService.GetComment(comment.ID); err != nil {
		response.FailWithMessage("查詢失敗", c)
	} else {
		response.OkWithData(gin.H{"recomment": recomment}, c)
	}
}
