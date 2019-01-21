package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Author     User
	Key        string // 为空时表示留言(message)，不为空时为对应博客的key值
	CommentKey string `gorm: "unique_index; not null"`
	Content    string
	AuthorID   uint
	Likes      int `gorm: "default : 0"`
}

func AddComment(author User, key, comment_key, content string, author_id uint) (comment *Comment, err error) {
	comment = &Comment{
		Author:     author,
		Key:        key,
		CommentKey: comment_key,
		Content:    content,
		AuthorID:   author_id,
	}
	err = db.Save(comment).Error

	return
}

func QueryCommentsWithPage(key string, page, limit int) (msgs []*Comment, err error) { //需要显示 评论或留言的 作者信息
	return msgs, db.Preload("Author").Where("Key = ?", key).Order("updated_at desc").Offset((page - 1) * limit).Limit(limit).Find(&msgs).Error //跳过前 (page-1)*limit 条 note，得到属于第一页的 note
}

func QueryCommentCount(key string) (count int, err error) { //留言或评论的总数, 当 key 为空时是留言
	return count, db.Model(&Comment{}).Where("Key = ?", key).Count(&count).Error
}

func DeleteCommentWithKey(key string) error {
	if err := db.Delete(&Comment{}, "Comment_Key = ?", key).Error; err != nil {
		return err
	}

	return nil
}
