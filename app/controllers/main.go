package controllers

import (
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/utils"
	//"runtime"
	//"strconv"
	//"strings"
	"time"
)

type MainController struct {
	BaseController
}

// 首页
func (this *MainController) Index() {
	this.Data["pageTitle"] = "系统概况"

	this.display()
}

// 获取系统时间
func (this *MainController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}
