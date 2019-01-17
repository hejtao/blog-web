package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
	"github.com/satori/go.uuid"
)

const (
	SESSION_USER_KEY = "SESSION_USER_KEY"
)

type NextPrepare interface {
	NextControllerPrepare()
}

type BaseController struct {
	beego.Controller
	User    models.User //登陆的用户
	IsLogin bool        //标识用户是否登陆
}

func (this *BaseController) Prepare() { //默认在其它 Controllers 之前执行
	this.Data["path"] = this.Ctx.Request.RequestURI // 将请求路径保存到 Path 变量里面

	this.IsLogin = false
	if u, ok := this.GetSession(SESSION_USER_KEY).(models.User); ok {
		this.User = u
		this.IsLogin = true
	}
	this.Data["user"] = this.User
	this.Data["isLogin"] = this.IsLogin

	if app, ok := this.AppController.(NextPrepare); ok { //？？？
		app.NextControllerPrepare()
	}
}

func (this *BaseController) Abort500(err error) {
	this.Data["error"] = err
	this.Abort("500") //调用 c.Error500(),其中 c 为 beego.ErrorController(&c)中的控制器
}

type StringMap map[string]interface{}

func (this *BaseController) ReturnJson(str_map StringMap) { //返回 json 响应到前端 [post]请求的返回
	this.Data["json"] = str_map

	this.ServeJSON()
}

func (this *BaseController) UUID() string {
	u, err := uuid.NewV4() //给每篇博客、评论生成一个id。每次调用生成新的id
	if err != nil {
		this.Abort500(my_errors.New("生成 UUID 时发生错误", err))
	}
	return u.String()
}
