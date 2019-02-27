package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type CommentController struct {
	BaseController
}

// func (this *CommentController) NextControllerPrepare() {
// 	if !this.IsLogin { //未登录
// 		this.Abort500(my_errors.NotLoginError{})
// 	}

// }

///comment_config
// @router /new/:key [get]
// func (this *CommentController) NewComment() {
// 	key := this.Ctx.Input.Param(":key")
// 	note, err := models.QueryNoteWithKey(key)
// 	if err != nil {
// 		this.Abort500(my_errors.New("评论时发生系统错误", err))
// 	}
// 	this.Data["note"] = note

// 	this.TplName = "new_comment.html"
// }

///comment_config
// @router /save/?:key [post]
func (this *CommentController) SaveComment() {
	if !this.IsLogin { //未登录
		this.Abort500(my_errors.NotLoginError{})
	}

	if this.User.Role == 10 { //未激活
		this.Abort500(my_errors.UnactivatedError{})
	}

	key := this.Ctx.Input.Param(":key")
	content := this.GetString("content", "")
	comment_key := this.UUID()

	cmt, err := models.AddComment(
		this.User,
		key,
		comment_key,
		content,
		this.User.ID,
	)
	if err != nil {
		this.Abort500(my_errors.New("评论或留言时发生系统错误", err))
	}

	if key == "" { //是留言
		this.ReturnJson(
			StringMap{
				"code":    3333,
				"message": cmt,
				"action":  "/message",
			},
		)
	} else {
		this.ReturnJson( //是评论
			StringMap{
				"code": 4444,
				//"action":  fmt.Sprintf("/note/%s", key),
				"comment": cmt,
			},
		)
	}

}

///comment_config
// @router /delete [get]
func (this *CommentController) DeleteComment() {
	if !this.IsLogin { //未登录
		this.Abort500(my_errors.NotLoginError{})
	}

	if this.User.Role == 10 { //未激活
		this.Abort500(my_errors.UnactivatedError{})
	}

	comment_key := this.GetString("comment_key", "")
	page, err := this.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1
	}

	err = models.DeleteCommentWithKey(comment_key)
	if err != nil {
		this.Abort500(my_errors.New("删除留言或评论时发生系统错误", nil))
	}

	url := fmt.Sprintf("/message?page=%d", page)
	this.Redirect(url, 302)

}

///comment_config
// @router /count/:comment_key/?:key [get]
func (this *CommentController) CommentCount() {
	comment_key := this.Ctx.Input.Param(":comment_key")
	key := this.Ctx.Input.Param(":key") //文章的key

	if comment_key != "placeholder" {
		err := models.DeleteCommentWithKey(comment_key)
		if err != nil {
			this.Abort500(my_errors.New("删除留言或评论时发生系统错误", nil))
		}
	}

	count, err := models.QueryCommentCount(key) //该文章的所有评论总数
	if err != nil {
		this.Abort500(my_errors.New("系统错误", err))
	}

	this.ReturnJson( //由于是 ajax请求必须返回json
		StringMap{
			"count": count,
		},
	)
}

///comment_config
// @router /query/?:key [get]
func (this *CommentController) CommentPageination() {
	key := this.Ctx.Input.Param(":key")

	page, err := this.GetInt("page", 1) //页码，默认 1
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := this.GetInt("limit", 10)
	if err != nil {
		page = 10
	}

	comments, err := models.QueryCommentsWithPage(key, page, limit)
	if err != nil {
		this.Abort500(my_errors.New("系统错误", err))
	}

	var dates []string
	for i := 0; i < len(comments); i++ {

		dates = append(
			dates,
			beego.Date(comments[i].CreatedAt, "m-d H:i"),
		)

	}

	user := this.User

	this.ReturnJson(
		StringMap{
			"code":     5555,
			"comments": comments,
			"user":     user,
			"is_login": this.IsLogin,
			"dates":    dates,
		},
	)
}
