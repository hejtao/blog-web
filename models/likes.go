package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type LikesInfo struct {
	gorm.Model
	Key      string
	AuthorID uint
}

func GetAndUpdateLikes(table, key string, author_id uint) (likes_value, code int, msg string, err error) {

	var info LikesInfo
	err = db.Model(&LikesInfo{}).Where("Key = ? and Author_ID = ?", key, author_id).Take(&info).Error
	if err == gorm.ErrRecordNotFound { //没有点赞记录
		info = LikesInfo{
			Key:      key,
			AuthorID: author_id,
		}

		if err = db.Save(&info).Error; err != nil {
			return 0, 0, "", err
		}

		if likes_value, err = updateLikes(table, key, 1); err != nil {
			return 0, 0, "", err
		}
		return likes_value, 1111, "点赞成功", nil

	} else if err != nil {
		return 0, 0, "", err
	}

	// if info.Flag { //点过赞了
	// 	if likes_value, err = GetAndUpdateLikes(table, key, -1); err != nil {
	// 		return 0, "", err
	// 	}
	// 	return likes_value, "你已取消点赞", nil
	// }

	if likes_value, err = updateLikes(table, key, 0); err != nil {
		fmt.Println("000000000000000000000000000000000000000000000")
		fmt.Println(table)
		fmt.Println(key)
		return 0, 0, "", err

	}
	return likes_value, 2222, "点过赞了", nil
}

//更新点赞次数
func updateLikes(table, key string, a int) (likes_value int, err error) {
	var temp struct {
		Likes int
	}

	if table == "notes" {
		err = db.Table(table).Where("Key = ?", key).Select("Likes").Scan(&temp).Error
		if err != nil {
			return 0, err
		}

		likes_value = temp.Likes + a

		err = db.Table(table).Where("Key = ?", key).Update("Likes", likes_value).Error
		if err != nil {
			return 0, err
		}

	}

	if table == "comments" {
		err = db.Table(table).Where("Comment_Key = ?", key).Select("Likes").Scan(&temp).Error
		if err != nil {
			return 0, err
		}

		likes_value = temp.Likes + a

		err = db.Table(table).Where("Comment_Key = ?", key).Update("Likes", likes_value).Error
		if err != nil {
			return 0, err
		}
	}

	return
}
