package service

import (
	"errors"
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

// DeleteComment 軟刪除
func (commentService *CommentService) DeleteComment(uuid string) (err error) {
	//db := global.Global_DB.Unscoped().Where("uuid=?", uuid).Delete(&model.Comment{}) // 永久刪除
	db := global.Global_DB.Where("uuid=?", uuid).Delete(&model.Comment{}) // 軟刪除

	if db.Error != nil || db.RowsAffected == 0 {
		return errors.New("刪除失敗或資料未異動")
	}

	return err
}

// UpdateComment 更新
func (commentService *CommentService) UpdateComment(uuid string, comment model.Comment) (err error) {
	db := global.Global_DB.Model(&comment).Where("uuid=?", uuid)
	db.Updates(map[string]interface{}{
		"uuid":           comment.Uuid,
		"parentid":       comment.Parentid,
		"comment":        comment.Comment,
		"author":         comment.Author,
		"update_comment": comment.UpdateComment,
		"favorite":       comment.Favorite, // bool更新必須要用map方式，否則會導致false不會更新
	})

	if db.Error != nil || db.RowsAffected == 0 {
		return errors.New("更新失敗或資料未異動")
	}

	return err
}

// GetComment 獲取
func (commentService *CommentService) GetComment(uuid string) (comment model.Comment, err error) {
	err = global.Global_DB.Where("uuid = ?", uuid).First(&comment).Error
	return
}
