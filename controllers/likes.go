package controllers

import (
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type LikesController struct {
	BaseController
}

func (this *LikesController) NextControllerPrepare() {
	if !this.IsLogin { //未登录
		this.Abort500(my_errors.NotLoginError{})
	}
}

///likes
// @router /:type/:key [post]
func (this *LikesController) Likes() {
	tp := this.Ctx.Input.Param(":type")
	key := this.Ctx.Input.Param(":key")

	table := "notes"
	switch tp {
	case "comment": //对评论点赞
		table = "comments"

	case "message": //对留言点赞
		table = "comments"

	case "note": //对文章点赞
		table = "notes"

	default:
		this.Abort500(my_errors.New("发生系统错误", nil))
	}

	likes_value, code, msg, err := models.GetAndUpdateLikes(table, key, this.User.ID)
	if err != nil {
		this.Abort500(my_errors.New("点赞时发生系统错误", err))
	}

	this.ReturnJson(
		StringMap{
			"code":        code,
			"msg":         msg,
			"likes_value": likes_value,
		},
	)
}
