package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Note struct {
	gorm.Model
	Author   User   // User 的 gorm.Model 被重写了
	Key      string `gorm:"unique; not null"`
	Title    string `gorm:"type : varchar(200)"`
	Content  string `gorm:"type : text"`
	Summary  string `gorm:"type : varchar(800)"`
	AuthorID uint
	Visits   int `gorm:"default: 0"`  //访问次数
	Likes    int `gorm:"default : 0"` //点赞次数

}

func AddNote(author User, key, title, content, summary string, author_id uint) error { //像数据库添加 note
	note_existed, err := QueryNoteWithKey(key)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			note := &Note{
				Author:   author,
				Key:      key,
				Title:    title,
				Content:  content,
				Summary:  summary,
				AuthorID: author_id,
			}

			if err2 := db.Save(note).Error; err2 != nil {
				return err2
			} else { // 更新成功
				return nil
			}
		}

		return err
	}

	if note_existed.Title != title { //标题是否修改
		note_existed.Title = title
		note_existed.CreatedAt = time.Now()
	}
	if note_existed.Content != content { //内容是否修改
		note_existed.Content = content
		note_existed.Summary = summary
		note_existed.CreatedAt = time.Now()
	}

	if err := db.Save(&note_existed).Error; err != nil {
		return err
	}

	return nil
}

func DeleteNoteWithKey(key string) error {
	if err := db.Delete(&Note{}, "Key = ?", key).Error; err != nil {
		return err
	}

	if err := db.Delete(&Comment{}, "Key = ?", key).Error; err != nil { //删除文章得评论
		return err
	}

	return nil
}

func QueryNoteWithKey(key string) (note Note, err error) {
	return note, db.Preload("Author").Where("Key = ?", key).Take(&note).Error
}

//获取文章并更新 Visits
func QueryNoteAndUpdateVisits(key string) (note Note, err error) {

	//若同一个用户查看该文章,则不更新 Visits         ?????????????????         未登录也可查看文章，所以没有意义
	// var info LikesAndVisitInfo
	// err = db.Model(&LikesAndVisitInfo{}).Where("Key = ? and Author_ID = ? and Is_Likes = false", key, visitor_id).Take(&info).Error

	// if err == gorm.ErrRecordNotFound { //该用户没有查看该文章的记录
	// 	info = LikesAndVisitInfo{
	// 		Key:      key,
	// 		AuthorID: visitor_id,
	// 		IsLikes:  false,
	// 	}
	// 	if err = db.Save(&info).Error; err != nil { //添加记录
	// 		return
	// 	}

	// 	new_value := note.Visits + 1
	// 	err = db.Model(&Note{}).Where("key = ?", key).Update("Visits", new_value).Error //更新Visits
	// 	if err != nil {
	// 		return
	// 	}

	// } else if err != nil {
	// 	return
	// }

	note, err = QueryNoteWithKey(key)
	if err != nil {
		return
	}

	new_value := note.Visits + 1
	err = db.Model(&Note{}).Where("key = ?", key).Update("Visits", new_value).Error //更新Visits
	if err != nil {                                                                 //虽然数据库更新了，但返回的 note 的 Visits 没有更新
		return
	}

	return note, nil
}

func QueryNotesWithPage(search string, page, limit int) (notes []*Note, err error) {
	return notes, db.Preload("Author").Where("Title like ?", fmt.Sprintf("%%%s%%", search)).Order("Created_At desc").Offset((page - 1) * limit).Limit(limit).Find(&notes).Error //跳过前 (page-1)*limit 条 note，得到属于第一页的 note
}

func QueryNoteCount(search string) (count int, err error) { //note 的总数
	return count, db.Model(&Note{}).Where("Title like ?", fmt.Sprintf("%%%s%%", search)).Count(&count).Error
}
