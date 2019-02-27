package main

import (
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jiangtaohe/blog-web/models"
	_ "github.com/jiangtaohe/blog-web/routers" //调用routers 的 init 方法
	"strings"
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
		func(a, b uint, c int, d bool) bool {
			if a == b || c == 0 && d {
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

	beego.AddFuncMap(
		"include",
		func(a, b interface{}) bool {
			if strings.Index(fmt.Sprintf("%v", b), fmt.Sprintf("%v", a)) == 0 {
				return true
			}
			return false
		},
	)

	//beego.AddFuncMap(
	//	"include",
	//	func(a, b string) bool {
	//		if strings.Index(b, a) == 0 {
	//			return true
	//		}
	//		return false
	//	},
	//)

	beego.AddFuncMap(
		"home",
		func(path interface{}) bool {
			if beego.Compare(path, "/") || strings.Index(fmt.Sprintf("%v", path), "/?") == 0 {
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
	beego.Run(":8080")
}
