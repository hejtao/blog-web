package controllers

import (
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type IndexController struct {
	BaseController
}

// var search string

//首页
// @router / [get]
func (this *IndexController) GetHome() {
	limit := 5                             //每页文章数量
	search := this.GetString("search", "") //从搜索框或者url提取搜索的关键字

	var page_count int = 0                      //分页总数
	count, err := models.QueryNoteCount(search) //文章总数
	if err != nil {
		this.Abort500(my_errors.New("显示主页时发生系统错误：", err))
	}

	if count%limit == 0 {
		page_count = count / limit
	} else {
		page_count = count/limit + 1
	}

	page, err := this.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}

	notes, err := models.QueryNotesWithPage(search, page, limit)
	if err != nil {
		this.Abort500(my_errors.New("显示主页时发生系统错误：", err))
	}

	this.Data["notes"] = notes
	this.Data["page_count"] = page_count
	this.Data["page"] = page
	this.Data["search"] = search

	this.TplName = "home2.html"
}

//留言
// @router /message [get]
func (this *IndexController) GetMessage() {
	limit := 6 //每页留言数量

	var page_count int = 0                     //分页总数
	count, err := models.QueryCommentCount("") //该文章的所有评论总数
	if err != nil {
		this.Abort500(my_errors.New("显示留言板时发生系统错误", err))
	}

	if count%limit == 0 {
		page_count = count / limit
	} else {
		page_count = count/limit + 1
	}

	page, err := this.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}

	comments, err := models.QueryCommentsWithPage("", page, limit)
	if err != nil {
		this.Abort500(my_errors.New("显示留言板时发生系统错误", err))
	}

	this.Data["comments"] = comments
	this.Data["user_id"] = this.User.ID
	this.Data["page_count"] = page_count
	this.Data["page"] = page

	this.TplName = "message3.html"
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
	}

	this.Data["content"] = "您的账号已激活, 请重新登陆"
	this.TplName = "error/500.html"
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

	this.TplName = "note2.html"
}
