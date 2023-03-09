package model

type CommentReq struct {
	Uuid     string `json:"uuid" form:"uuid" uri:"uuid" `
	Parentid string `json:"parentid" form:"parentid" binding:"required"`
	Comment  string `json:"comment" form:"comment" binding:"required"`
	Author   string `json:"author" form:"author" binding:"required"`
	Update   string `json:"update" form:"update" binding:"required"`
	Favorite *bool  `json:"favorite" form:"favorite" binding:"required"`
}

type UuidCommentReq struct {
	Uuid string `uri:"uuid"`
}

type UpdateCommentReq struct {
	Uuid     string `json:"uuid" form:"uuid" uri:"uuid" binding:"required"`
	Parentid string `json:"parentid" form:"parentid" binding:"required"`
	Comment  string `json:"comment" form:"comment" binding:"required"`
	Author   string `json:"author" form:"author" binding:"required"`
	Update   string `json:"update" form:"update" binding:"required"`
	Favorite *bool  `json:"favorite" form:"favorite" binding:"required"`
}
