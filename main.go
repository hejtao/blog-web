package main

import (
	"encoding/gob"
	"github.com/astaxie/beego"
	"github.com/jiangtaohe/blog-web/models"
	_ "github.com/jiangtaohe/blog-web/routers" //调用routers 的 init 方法
)

func initTemplate() {
	beego.AddFuncMap(
		"add",
		func(a, b int) int {
			return a + b
		},
	)

	beego.AddFuncMap(
		"fix_bar",
		func(a, b uint, c bool) bool {
			if a == b && c {
				return true
			}

			return false
		},
	)

	beego.AddFuncMap(
		"and",
		func(a, b bool) bool {
			if a && b {
				return true
			}
			return false
		},
	)
}

func initSession() {
	gob.Register(models.User{})                      //beego的session序列号是用gob方式
	beego.BConfig.WebConfig.Session.SessionOn = true //开启 session
	beego.BConfig.WebConfig.Session.SessionName = "myblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}

func main() {
	initSession()
	initTemplate()
	beego.Run()
}
