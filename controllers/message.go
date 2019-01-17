package controllers

import (
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type MessageController struct {
	BaseController
}

// func (this *MessageController) NextControllerPrepare() {
// 	if !this.IsLogin { //未登录
// 		this.Abort500(my_errors.NotLoginError{})
// 	}
// }

///message_config
// @router /count [get]
func (this *MessageController) MessageCount() {
	count, err := models.QueryCommentCount("")
	if err != nil {
		this.Abort500(my_errors.New("显示留言板时发生系统错误", err))
	}

	this.ReturnJson( //由于是 ajax请求必须返回json
		StringMap{
			"count": count,
		},
	)
}

//message_config
// @router /query [get]
func (this *MessageController) MessagePageination() {
	page, err := this.GetInt("page", 1) //页码，默认 1
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := this.GetInt("limit", 10) //每页留言数量， 默认 10
	if err != nil {
		page = 10
	}

	msgs, err := models.QueryCommentsWithPage("", page, limit)
	if err != nil {
		this.Abort500(my_errors.New("显示留言板时发生系统错误", err))
	}

	this.ReturnJson(
		StringMap{
			"code":     5555,
			"messages": msgs,
		},
	)
}
