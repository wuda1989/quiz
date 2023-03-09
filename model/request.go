package model

type CommentReq struct {
	Uuid     string `json:"uuid" form:"uuid" `
	Parentid string `json:"parentid" form:"parentid" binding:"required"`
	Comment  string `json:"comment" form:"comment" binding:"required"`
	Author   string `json:"author" form:"author" binding:"required"`
	Update   string `json:"update" form:"update" binding:"required"`
	Favorite *bool  `json:"favorite" form:"favorite" binding:"required"`
}
