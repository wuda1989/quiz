package quiz

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"quiz/model"
	response "quiz/model"
	"quiz/service"
)

type CommentApi struct {
}

var commentService = service.ServiceGroupApp.CommentService

func (commentApi *CommentApi) CreateComment(c *gin.Context) {
	var params model.CommentReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	newUUID := uuid.New().String()
	comment := model.Comment{
		Uuid:          newUUID,
		Parentid:      params.Parentid,
		Comment:       params.Comment,
		Author:        params.Author,
		UpdateComment: params.Update,
		Favorite:      *params.Favorite,
	}

	if err = commentService.CreateComment(comment); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithJson(comment, c)
}

// FindComment 用UUID查詢
func (commentApi *CommentApi) FindComment(c *gin.Context) {
	var params model.UuidCommentReq
	err := c.ShouldBindUri(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if recomment, err := commentService.GetComment(params.Uuid); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithJson(recomment, c)
	}
}

func (commentApi *CommentApi) UpdateComment(c *gin.Context) {
	var paramsUri model.UuidCommentReq
	err := c.ShouldBindUri(&paramsUri)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var params model.UpdateCommentReq
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	comment := model.Comment{
		Uuid:          params.Uuid,
		Parentid:      params.Parentid,
		Comment:       params.Comment,
		Author:        params.Author,
		UpdateComment: params.Update,
		Favorite:      *params.Favorite,
	}

	if err = commentService.UpdateComment(paramsUri.Uuid, comment); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithJson(comment, c)

}

func (commentApi *CommentApi) DeleteComment(c *gin.Context) {
	var paramsUri model.UuidCommentReq
	err := c.ShouldBindUri(&paramsUri)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err = commentService.DeleteComment(paramsUri.Uuid); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("刪除成功", c)
}
