package model

import "quiz/global"

// Comment ORM
type Comment struct {
	global.IdModel
	Uuid     string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:UUID;size:191;"`
	Parentid string `json:"parentid" form:"parentid" gorm:"column:parentid;comment:parentID;size:191;"`
	Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:內容;size:191;"`
	Author   string `json:"author" form:"author" gorm:"column:author;comment:來源;size:191;"`
	Favorite bool   `json:"favorite" form:"favorite" gorm:"column:favorite;comment:最愛;"`
}

// TableName Comment
func (Comment) TableName() string {
	return "comment"
}
