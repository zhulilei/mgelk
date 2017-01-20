package libs

import ()

type Result struct {
	Data    interface{}
	Code    int
	Message string
}

type Json struct {
	SearchSourceJSON interface{} `json:"searchSourceJSON"`
}

var NginxUrl = "http://10.170.164.76:9200/.kibana/visualization/"

var NginxDash = "http://10.170.164.76:9200/.kibana/dashboard/"
