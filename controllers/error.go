package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/jiangtaohe/blog-web/my_errors"
)

type ErrorController struct {
	BaseController
}

// ajax: {code:, msg:, origin:erorr}
func (this *ErrorController) Error404() {
	if this.IsAjax() { //是ajax请求
		this.errorJson(my_errors.Error404{}) //错误信息以json格式发送到前端
	}

	this.TplName = "error/404.html"
}

func (this *ErrorController) Error500() {

	//error 包含了 Error (Error比error更细分，实现了更多的方法)
	err, ok := this.Data["error"].(error) //是 error 吗
	if !ok {                              //不是则转成 Error
		err = my_errors.New("系统发生未知错误", nil) //err要么是 error，要么是 Error
	}

	err2, ok := err.(my_errors.Error) //是 Error 吗
	if !ok {                          //不是 Error，那么说明是 error
		err2 = my_errors.New(err.Error(), nil)
	}

	if err2.GetOrigin() != nil {
		logs.Info(err2, ": ", err2.GetOrigin()) //默认调用了 Error() 方法,等价于 logs.Info(err2.Error(), err2.GetOrigin().Error())
	}

	if this.IsAjax() { //如果是 ajax请求 返回 json 数据就 ok 了，否则就跳转到 this.TplName 指定的页面
		this.errorJson(err2)
	} else {
		this.Data["content"] = err2.Error()
	}

	this.TplName = "error/500.html"
}

func (this *ErrorController) errorJson(err my_errors.Error) {
	this.Ctx.Output.Status = 200
	this.ReturnJson(
		StringMap{
			"code": err.ErrorCode(),
			"msg":  err.Error(),
		},
	)
}
