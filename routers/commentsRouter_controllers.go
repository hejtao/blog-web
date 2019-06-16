package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"],
        beego.ControllerComments{
            Method: "CommentCount",
            Router: `/count/:comment_key/?:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"],
        beego.ControllerComments{
            Method: "DeleteComment",
            Router: `/delete`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"],
        beego.ControllerComments{
            Method: "CommentPageination",
            Router: `/query/?:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:CommentController"],
        beego.ControllerComments{
            Method: "SaveComment",
            Router: `/save/?:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetHome",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetAbout",
            Router: `/about`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Activation",
            Router: `/activation/:md5`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "GetMessage",
            Router: `/message`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "NoteDetail",
            Router: `/note/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Regislation",
            Router: `/reg`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:LikesController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:LikesController"],
        beego.ControllerComments{
            Method: "Likes",
            Router: `/:type/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"],
        beego.ControllerComments{
            Method: "DeleteNote",
            Router: `/delete/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"],
        beego.ControllerComments{
            Method: "EditNote",
            Router: `/edit/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"],
        beego.ControllerComments{
            Method: "NewNote",
            Router: `/new`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:NoteController"],
        beego.ControllerComments{
            Method: "SaveNote",
            Router: `/save/:key`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/jiangtaohe/blog-web/controllers:UserController"],
        beego.ControllerComments{
            Method: "Regislation",
            Router: `/reg`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
