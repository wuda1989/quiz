package service

import (
	"quiz/global"
	"quiz/model"
)

type CommentService struct {
}

// CreateComment 創建
func (commentService *CommentService) CreateComment(comment model.Comment) (err error) {
	err = global.Global_DB.Create(&comment).Error
	return err
}

// DeleteComment 刪除
func (commentService *CommentService) DeleteComment(comment model.Comment) (err error) {
	err = global.Global_DB.Delete(&comment).Error
	return err
}

// UpdateComment 更新
func (commentService *CommentService) UpdateComment(comment model.Comment) (err error) {
	err = global.Global_DB.Save(&comment).Error
	return err
}

// GetComment 獲取
func (commentService *CommentService) GetComment(uuid uint) (comment model.Comment, err error) {
	err = global.Global_DB.Where("uuid = ?", uuid).First(&comment).Error
	return
}
