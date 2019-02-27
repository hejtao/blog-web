package routers

import (
	"github.com/jiangtaohe/my-web/my-web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
