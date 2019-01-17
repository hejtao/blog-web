package controllers

import (
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type CommentController struct {
	BaseController
}

var key string

func (this *CommentController) NextControllerPrepare() {
	if !this.IsLogin { //未登录
		this.Abort500(my_errors.NotLoginError{})
	}

	if str, ok := this.Data["key"].(string); ok { //来自 index.go NoteDetail()
		key = str

	}

}

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
// @router /count/?:key [get]
func (this *CommentController) CommentCount() {
	key := this.Ctx.Input.Param(":key")
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

	this.ReturnJson(
		StringMap{
			"code":     5555,
			"comments": comments,
		},
	)
}
