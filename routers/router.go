package routers

import (
	"github.com/astaxie/beego"
	"github.com/jiangtaohe/blog-web/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{},
		&controllers.BaseController{},
		&controllers.UserController{},
	)

	beego.AddNamespace(
		beego.NewNamespace("note_config", beego.NSInclude(&controllers.NoteController{})),
		beego.NewNamespace("likes", beego.NSInclude(&controllers.LikesController{})),
		beego.NewNamespace("comment_config", beego.NSInclude(&controllers.CommentController{})),
	)

	//beego.Router("/note/new", &controllers.NoteController{}, "get:Write")
	// beego.SetStaticPath("/note_config/static", "static")
}
