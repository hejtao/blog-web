package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/utils"
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
	"math/rand"
	"regexp"
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
	//checkbox := this.GetString("checkbox", "")

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
	match, err := checkEmail(email)
	if err != nil {
		this.Abort500(my_errors.New("", err))
	}
	if !match {
		this.Abort500(my_errors.New("你输入的邮箱地址格式不正确, 请重新输入", err))
	}

	pwd := this.GetString("password", "")
	pwd2 := this.GetString("password2", "") //确认密码
	if strings.Compare(pwd, pwd2) != 0 {
		this.Abort500(my_errors.New("密码不一致, 请重新输入", nil))
	}

	uuid := this.UUID() //由邮箱和uuid生成md5
	raw_str := email + uuid
	md5, err := creatMD5(raw_str)
	if err != nil {
		this.Abort500(my_errors.New("", err))
	}

	if err := this.sendEmail(name, email, md5); err != nil { //发送激活邮件
		this.Abort500(my_errors.New("", err))
	}

	avatar := creatRandAvatar() //随机生成用户图像

	if err := models.AddUser(name, email, pwd, avatar, md5, 10); err != nil { //将用户添加到数据库,未激活状态
		this.Abort500(my_errors.New("", err))
	}

	this.ReturnJson(
		StringMap{
			"code":   "AAAA", //当前端收到 code 为 "AAAA" 时，提示注册成功,并跳转到登陆界面
			"action": "/login",
		},
	)
}

func checkEmail(email string) (match bool, err error) {
	reg := `^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z0-9]{2,6}$`
	match, err = regexp.MatchString(reg, email)

	return
}

func (this *UserController) sendEmail(name, email, md5 string) error {
	date := beego.Date(time.Now(), "Y-m-d H:i:s")

	config := `{"username":"896419134@qq.com","password":"ihlwdszyhbhebfhi","host":"smtp.qq.com","port":587}`
	smtp := utils.NewEMail(config)
	smtp.To = []string{email}
	smtp.From = "896419134@qq.com"
	smtp.Subject = "Meetoo账户激活"
	smtp.HTML = fmt.Sprintf(`<!DOCTYPE html>
								<html lang="">
								    <meta charset="UTF-8">                                      
								        <table width="33%%" border="0" align="center" cellpadding="20" cellspacing="10" >
								                <tbody>
								                        <td style="background-color: #f2f2f2;">
								                                <h4>亲爱的%s，您好：</h4>
								                                <p>您于 %s 在 Meetoo 完成了注册，请您点击下面的链接激活账户，</p>
								                                
								                                <p><a href="http://localhost/activation/%s">http://localhost/activation/%s</a></p>

								                                <p>如果您没有在 Meetoo 注册，请忽略该邮件。</p>

								                                <p> Meetoo ! (*^_^*) </p>  　　　　　　　　　　　　　　　　　　　
								                        </td>
								                </tbody>
								        </table>                                                                                                   
								</html>`, name, date, md5, md5)
	// smtp.AttachFile("1.jpg") // 附件

	if err := smtp.Send(); err != nil {
		return err
	}

	return nil
}

func creatMD5(raw_str string) (md5 string, err error) {
	url := fmt.Sprintf("https://devops-api.com/api/v1/md5?rawstr=%s", raw_str)
	req := httplib.Get(url)
	req.Header("DEVOPS-API-TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoicHVibGljIiwidXBkYXRlVGltZSI6MTUzNTUzMzQ4NH0.JKxOjbtkmZC9FpPPkmF6u6AzBEYJt6m-yYyYr9wmx18") //设置 API 请求头

	res, err := req.String()
	if err != nil {
		return
	}

	md5 = res[40+len(raw_str) : 72+len(raw_str)]
	return
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
