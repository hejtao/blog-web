package controllers

import (
	"fmt"
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
	"math/rand"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (this *UserController) Login() {
	email := this.GetString("email", "")  //<input ... name="email" required lay-verify="required" ... >邮箱不能为空
	pwd := this.GetString("password", "") //<input ... name="password" required lay-verify="required" ...> 密码不能为空
	checkbox := this.GetString("checkbox", "")
	fmt.Println("11111111111111111111111111111111111111111111111111111")
	fmt.Println(checkbox)
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

//以游客身份登陆
// @router /visitor [get]
// func (this *UserController) Visitor() {
// 	user, err := models.QueryWithEmailAndPwd("", "") //查询数据库
// 	if err != nil {
// 		this.Abort500(my_errors.New("以游客登入时发生系统错误", err))
// 	}

// 	this.SetSession(SESSION_USER_KEY, user)
// 	this.TplName = "home.html"
// }

// @router /reg [post]
func (this *UserController) Regislation() {
	name := this.GetString("name", "")
	if models.HasNameExisted(name) {
		this.Abort500(my_errors.New("您的昵称已被占用", nil))
	}

	email := this.GetString("email", "")
	if models.HasEmailExisted(email) {
		this.Abort500(my_errors.New("您的邮箱已被注册", nil))
	}

	pwd := this.GetString("password", "")
	pwd2 := this.GetString("password2", "") //确认密码
	if strings.Compare(pwd, pwd2) != 0 {
		this.Abort500(my_errors.New("密码不一致, 请重新输入", nil))
	}

	avatar := creatRandAvatar()

	if err := models.AddUser(name, email, pwd, avatar, 1); err != nil { //将用户添加到数据库
		this.Abort500(my_errors.New("注册系统错误", err))
	}

	this.ReturnJson(
		StringMap{
			"code":   "AAAA", //当前端收到 code 为 "AAAA" 时，提示注册成功,并跳转到登陆界面
			"action": "/login",
		},
	)
}

func creatRandAvatar() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("/static/images/%d.png", r.Intn(61))
}

// @router /logout [get]
func (this *UserController) Logout() {

	this.DelSession(SESSION_USER_KEY)
	//this.IsLogin = false
	this.Redirect("/", 302)
}
