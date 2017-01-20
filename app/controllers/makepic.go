package controllers

import (
	"github.com/astaxie/beego"
	"mgelk/app/libs"
	"mgelk/app/models"
	"strings"
	//"time"
)

type MakePicController struct {
	BaseController
}

// 任务列表
func (this *MakePicController) List() {
	this.Data["pageTitle"] = "选择集群"
	//this.Data["pageBar"] = libs.NewPager(page, int(count), this.pageSize, beego.URLFor("TaskController.List", "groupid", groupId), true).ToString()

	clusterId, _ := this.GetInt("clusterid")

	clusters := models.GetCluster()
	beego.Critical(clusters)

	this.Data["clusterid"] = clusterId
	this.Data["clusters"] = clusters
	this.display()

}

func (this *MakePicController) Add() {
	clusterId, _ := this.GetInt("clusterid")
	switch clusterId {
	case 1:
		this.Redirect("/index", 302)
		return
	case 2:
		this.AddNginx()
		return
	}
}

func (this *MakePicController) AddNginx() {
	this.actionName = "addnginx"
	if this.isPost() {
		index := strings.TrimSpace(this.GetString("index_name"))
		domain := strings.TrimSpace(this.GetString("server_name"))
		fromtime := strings.TrimSpace(this.GetString("from_time"))
		totime := strings.TrimSpace(this.GetString("to_time"))
		topX, _ := this.GetInt("topN")

		if baseCode := libs.NginxtempBase(index, domain, topX); baseCode != 201 {
			this.ajaxMsg("make-base-error", MSG_ERR)
		} else {
			domainUri := libs.Geturi(index, domain, fromtime, totime, topX)
			//fmt.Println("====")
			//fmt.Println(domainUri)
			//fmt.Println("====")
			for i, v := range domainUri[domain] {
				//fmt.Println(v)
				if uricode := libs.NginxtempUri(index, domain, v, i+1); uricode != 201 {
					this.ajaxMsg("make-nginx-topx-error", MSG_ERR)
				}
			}
		}

		if dashCode := libs.Nginxdashboard(domain, topX); dashCode != 201 {
			this.ajaxMsg("make-nginx-dashboard-error", MSG_ERR)
		}
		this.ajaxMsg("success created", MSG_OK)
	}

	this.Data["pageTitle"] = "nginx topN URL 视图"
	this.display()
}
