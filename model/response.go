package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type CommentResp struct {
	Uuid     string `json:"uuid"`
	Parentid string `json:"parentid"`
	Comment  string `json:"comment"`
	Author   string `json:"author"`
	Update   string `json:"update"`
	Favorite bool   `json:"favorite"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	if code == ERROR {
		c.JSON(http.StatusBadRequest, Response{
			code,
			data,
			msg,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithJson(comment Comment, c *gin.Context) {
	commentResp := CommentResp{
		Uuid:     comment.Uuid,
		Parentid: comment.Parentid,
		Comment:  comment.Comment,
		Author:   comment.Author,
		Update:   comment.UpdateComment,
		Favorite: comment.Favorite,
	}
	c.JSON(http.StatusOK, commentResp)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "查詢成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失敗", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
