package routers

import (
	"beegolearn/src/mgelk/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/zhuapi/v1/nginx/template", &controllers.HomeController{})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Println(err)
	}
}
