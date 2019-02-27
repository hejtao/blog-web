package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //数据库驱动
	"os"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name    string `gorm: "unique_index"`
	Email   string `gorm: "unique_index"`
	Avatar  string
	Pwd     string //密码
	ActCode string //账户的激活码
	Role    int    //0 管理员 1已激活用户 10未激活用户
}

func init() {
	var err error

	if err = os.MkdirAll("data", 0777); err != nil { //在上级目录创建 data 目录，存在不创建
		panic(err.Error())
	}

	db, err = gorm.Open("sqlite3", "data/data.db") //在 data 目录下生成数据库文件 data.db
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	db.AutoMigrate(&User{}, &Note{}, &Comment{}, &LikesInfo{})

	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(
			&User{
				Name:   "米琴香光",
				Email:  "hejtao@outlook.com",
				Pwd:    "123",
				Avatar: "/static/images/admin.png",
				Role:   0,
			},
		) //初始化一个管理员
	}
}

func QueryWithEmailAndPwd(email, pwd string) (user User, err error) {
	return user, db.Where("Email = ? and Pwd = ?", email, pwd).Take(&user).Error
}

func QueryAndActivate(act_code string) error {
	var user User
	if err := db.Where("Act_Code = ?", act_code).Take(&user).Error; err != nil {
		return err
	}

	if err := db.Model(&User{}).Where("Act_Code = ?", act_code).Update("Role", 1).Error; err != nil {
		return err
	}

	return nil
}

func HasNameExisted(name string) bool {
	var user User
	return db.Where("Name = ?", name).Take(&user).Error == nil //查询无误，存在
}

func HasEmailExisted(email string) bool {
	var user User
	return db.Where("Email = ?", email).Take(&user).Error == nil
}

func AddUser(name, email, pwd, avatar, act_code string, role int) error {
	user := &User{
		Name:    name,
		Email:   email,
		Pwd:     pwd,
		Avatar:  avatar,
		ActCode: act_code,
		Role:    role,
	}
	return db.Save(user).Error
}
