package controllers

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/jiangtaohe/blog-web/models"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) NextControllerPrepare() {
	if !this.IsLogin { //未登录
		this.Abort500(my_errors.NotLoginError{})
	}

	if this.User.Role == 10 { //是游客
		this.Abort500(my_errors.UnactivatedError{})
	}
}

///note_config
// @router /new [get]
func (this *NoteController) NewNote() { //写博客按钮触发
	var note models.Note
	note.Key = this.UUID()
	this.Data["note"] = note

	this.TplName = "new_note.html"
}

///note_config
// @router /edit/:key [get]
func (this *NoteController) EditNote() { //修改博客按钮触发
	key := this.Ctx.Input.Param(":key")
	note, err := models.QueryNoteWithKey(key)
	if err != nil {
		this.Abort500(my_errors.New("编辑文章时发生系统错误", err))
	}

	this.Data["note"] = note

	this.TplName = "new_note.html"
}

///note_config
// @router /save/:key [post]
func (this *NoteController) SaveNote() { //提交按钮触发
	key := this.Ctx.Input.Param(":key")
	title := this.GetString("title", "")
	raw_content := this.GetString("content", "") //有可能是 html 文本
	content, err := htmlFilter(raw_content)
	if err != nil {
		this.Abort500(my_errors.New("保存文章时发生系统错误", err))
	}

	if len(content) == 0 {
		this.ReturnJson(
			StringMap{
				"code": 0,
				"msg":  "正文不能为空",
			},
		)
		return
	}

	summary := getSummary(content)

	err = models.AddNote(
		this.User,
		key,
		title,
		raw_content,
		summary,
		this.User.ID,
	)

	if err != nil {
		this.Abort500(my_errors.New("保存文章时发生系统错误：", err))
	}

	this.ReturnJson(
		StringMap{
			"code":   7777,
			"action": fmt.Sprintf("/note/%s", key),
		},
	)
}

///note_config
// @router /delete/:key [post]
func (this *NoteController) DeleteNote() { //由删除固定按钮触发
	key := this.Ctx.Input.Param(":key")
	if err := models.DeleteNoteWithKey(key); err != nil {
		this.ReturnJson(
			StringMap{
				"code":   0, //0 表示删除失败，返回错误信息并停留在当前页面
				"msg":    err.Error(),
				"action": fmt.Sprintf("/note/%s", key),
			},
		)
	} else {
		this.ReturnJson(
			StringMap{
				"code":   6666,
				"action": "/",
			},
		)
	}
}

func htmlFilter(html string) (string, error) {
	var bytes_buf bytes.Buffer

	if _, err := bytes_buf.Write([]byte(html)); err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(&bytes_buf) //过滤 html 文本
	if err != nil {
		return "", err
	}

	txt := doc.Find("body").Text()

	return txt, nil
}

func getSummary(content string) string { //从正文中提取摘要

	if len(content) > 500 {
		content = content[:500]
	}

	return content
}
