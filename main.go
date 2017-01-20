package main

import (
	"github.com/astaxie/beego"
	"html/template"
	"mgelk/app/controllers"
	"net/http"
)

const VERSION = "1.0.0"

func main() {

	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	//首页
	beego.Router("/index", &controllers.MainController{}, "*:Index")
	//获取时间
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	//登陆页
	//beego.Router("/login", &controllers.MainController{}, "*:Login")
	//登出页
	//beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	//后台页
	//beego.Router("/backend", &controllers.BackController{}, "*:Backend")
	//帮助页
	//beego.Router("/help", &controllers.HelpController{}, "*:Index")
	//视图页
	beego.AutoRouter(&controllers.MakePicController{})
	//报警页
	//f///////vvvvvvvvvvvvvvvvvvvvvvvvvvvv beego.AutoRouter(&controllers.MakeAlarmController{})

	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.Run()

}
