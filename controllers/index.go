package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type IndexController struct {
	BaseController
}

var search string

//首页
// @router / [get]
func (this *IndexController) GetHome() {
	search = this.GetString("search", "") //从搜索框或者url提取搜索的关键字,更新到全局变量 search
	this.TplName = "home.html"            //在 views 目录下查找 home.html 文件
}

// @router /count [get]
func (this *IndexController) NoteCount() { //进入 home.html 后, <div id="test2" class="paging"></div> 触发 blog.js 的主页分页代码 if ($("#test2").size()>0)
	count, err := models.QueryNoteCount(search)
	if err != nil {
		this.Abort500(my_errors.New("显示主页时发生系统错误：", err))
	}

	this.ReturnJson( //由于是 ajax请求必须返回json
		StringMap{
			"count":  count,
			"search": search,
		},
	)
}

// @router /query [get]
func (this *IndexController) HomePageination() {
	page, err := this.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := this.GetInt("limit", 5)
	if err != nil {
		page = 5
	}

	notes, err := models.QueryNotesWithPage(search, page, limit)
	if err != nil {
		this.Abort500(my_errors.New("显示主页时发生系统错误：", err))
	}

	var dates []string
	for i := 0; i < len(notes); i++ {

		dates = append(
			dates,
			beego.Date(notes[i].CreatedAt, "Y-m-d H:i:s"),
		)

	}

	this.ReturnJson(
		StringMap{
			"code":  9999,
			"notes": notes,
			"dates": dates,
		},
	)
}

//留言
// @router /message [get]
func (this *IndexController) GetMessage() {
	this.TplName = "message.html" // [get] 请求的返回
}

//关于
// @router /about [get]
func (this *IndexController) GetAbout() {
	this.TplName = "about.html"
}

// @router /login [get]
func (this *IndexController) Login() {
	// this.TplName = "login.html"

	this.TplName = "login.html"
}

// @router /reg [get]
func (this *IndexController) Regislation() {
	this.TplName = "reg.html"
}

// @router /activation/:md5 [get]
func (this *IndexController) Activation() {

	md5 := this.Ctx.Input.Param(":md5")
	if err := models.QueryAndActivate(md5); err != nil {
		this.Abort500(my_errors.New("激活账号时发生系统错误", err))
	} else {
		this.Data["content"] = "您的账号已激活, 请重新登陆"
		this.TplName = "error/500.html"
	}
}

// @router /note/:key [get]
func (this *IndexController) NoteDetail() {
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteAndUpdateVisits(key)
	if err != nil {
		this.Abort500(my_errors.New("显示文章时发生系统错误", err))
	}
	this.Data["note"] = note // 发给 note.html

	// comments, err := models.QueryCommentWithKey(key) //该文章的所有评论
	// if err != nil {
	// 	this.Abort500(my_errors.New("显示文章时发生系统错误", err))
	// }
	// this.Data["comments"] = comments

	this.TplName = "note.html"
}
