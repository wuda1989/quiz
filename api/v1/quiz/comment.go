package quiz

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"quiz/model"
	response "quiz/model"
	"quiz/service"
	"time"
)

type CommentApi struct {
}

var commentService = service.ServiceGroupApp.CommentService

type CommentResp struct {
	Uuid     string `json:"uuid"`
	Parentid string `json:"parentid"`
	Comment  string `json:"comment"`
	Author   string `json:"author"`
	Update   string `json:"update"`
	Favorite bool   `json:"favorite"`
}

func (commentApi *CommentApi) CreateComment(c *gin.Context) {
	var params model.CommentReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.Fail(c)
		return
	}

	newUUID := uuid.New().String()
	comment := model.Comment{
		Uuid:     newUUID,
		Parentid: params.Parentid,
		Comment:  params.Comment,
		Author:   params.Author,
		Favorite: *params.Favorite,
	}

	if err := commentService.CreateComment(comment); err != nil {
		response.FailWithMessage("create comment fail", c)
		return
	}

	now := time.Now()
	commentResp := CommentResp{
		Uuid:     comment.Uuid,
		Parentid: comment.Parentid,
		Comment:  comment.Comment,
		Author:   comment.Author,
		Update:   now.Format("2006-01-02T15:04:05Z"),
		Favorite: comment.Favorite,
	}

	response.OkWithJson(commentResp, c)
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
