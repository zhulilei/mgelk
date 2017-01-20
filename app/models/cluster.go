package models

import (
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"net/url"
)

type Cluster struct {
	Id          int
	ClusterName string
}

func GetCluster() (ClusterList []Cluster) {
	ClusterList = append(ClusterList, Cluster{Id: 1, ClusterName: "syslog"})
	ClusterList = append(ClusterList, Cluster{Id: 2, ClusterName: "nginx"})
	return ClusterList
}
