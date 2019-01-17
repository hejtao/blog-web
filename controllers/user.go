package controllers

import (
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
	"strings"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	email := this.GetString("email", "")  //<input ... name="email" required lay-verify="required" ... >邮箱不能为空
	pwd := this.GetString("password", "") //<input ... name="password" required lay-verify="required" ...> 密码不能为空

	user, err := models.QueryWithEmailAndPwd(email, pwd) //查询数据库
	if err != nil {
		this.Abort500(my_errors.New("邮箱或密码错误", err))
	}

	this.SetSession(SESSION_USER_KEY, user) //用户登陆成功，则保存到 session
	this.ReturnJson(
		StringMap{
			"code":   8888, //当前端收到 code 为 8888 时，提示登陆成功,并跳转到首页
			"action": "/",
		},
	)
}

// @router /reg [post]
func (this *UserController) Register() {
	name := this.GetString("name", "")
	email := this.GetString("email", "")
	pwd := this.GetString("password", "")
	pwd2 := this.GetString("password2", "") //确认密码

	if strings.Compare(pwd, pwd2) != 0 {
		this.Abort500(my_errors.New("密码不一致, 请重新输入", nil))
	}

	if models.HasNameExisted(name) {
		this.Abort500(my_errors.New("您的昵称已被占用", nil))
	}

	if models.HasEmailExisted(email) {
		this.Abort500(my_errors.New("您的邮箱已被注册", nil))
	}

	if err := models.AddUser(name, email, pwd, "static/images/anonimity.png", 1); err != nil { //将用户添加到数据库
		this.Abort500(my_errors.New("注册系统错误", err))
	}

	this.ReturnJson(
		StringMap{
			"code":   10, //当前端收到 code 为 10 时，提示注册成功,并跳转到登陆界面
			"action": "/login",
		},
	)
}

// @router /logout [get]
func (this *UserController) Logout() {

	this.DelSession(SESSION_USER_KEY)
	//this.IsLogin = false
	this.Redirect("/", 302)
}
