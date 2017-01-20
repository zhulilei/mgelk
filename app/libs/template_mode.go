package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func NginxtempBase(index, domain string, topX int) (code int) {
	url := NginxUrl + domain + "-" + "base"
	jsonstr := fmt.Sprintf("{\"index\":\"%s\",\"query\":{\"query_string\":{\"query\":\"host: \\\"%s\\\"\",\"analyze_wildcard\":true}},\"filter\":[]}", index, domain)
	visStatetmp := fmt.Sprintf("{\"title\":\"New Visualization\",\"type\":\"pie\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":true,\"isDonut\":false},\"aggs\":[{\"id\":\"1\",\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"uri.raw\",\"size\":%d,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}", topX)
	searchjson := Json{SearchSourceJSON: jsonstr}
	m := map[string]interface{}{
		//"visState":              "{\"title\":\"New Visualization\",\"type\":\"pie\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":true,\"isDonut\":false},\"aggs\":[{\"id\":\"1\",\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"uri.raw\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}",
		"visState":              visStatetmp,
		"title":                 domain + "-" + "base",
		"version":               1,
		"kibanaSavedObjectMeta": searchjson,
	}
	mJson, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", url, contentReader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	return resp.StatusCode
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

func NginxtempUri(index, domain, uri string, i int) (code int) {
	id := strconv.Itoa(i)
	url := NginxUrl + domain + "-" + "uri" + id
	jsonstr := fmt.Sprintf("{\"index\":\"%s\",\"query\":{\"query_string\":{\"query\":\"host: \\\"%s\\\" && uri: \\\"%s\\\"\",\"analyze_wildcard\":true}},\"filter\":[]}", index, domain, uri)
	//visStatetmp := fmt.Sprintf("{\"title\":\"New Visualization\",\"type\":\"pie\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":true,\"isDonut\":false},\"aggs\":[{\"id\":\"1\",\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"method.raw\",\"size\":%d,\"order\":\"desc\",\"orderBy\":\"1\"}},{\"id\":\"3\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"status.raw\",\"size\":%d,\"order\":\"desc\",\"orderBy\":\"1\"}},{\"id\":\"4\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"geoip.city_name.raw\",\"size\":%d,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}", topX)
	searchjson := Json{SearchSourceJSON: jsonstr}
	m := map[string]interface{}{
		"visState": "{\"title\":\"New Visualization\",\"type\":\"pie\",\"params\":{\"shareYAxis\":true,\"addTooltip\":true,\"addLegend\":true,\"isDonut\":false},\"aggs\":[{\"id\":\"1\",\"type\":\"count\",\"schema\":\"metric\",\"params\":{}},{\"id\":\"2\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"method.raw\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\"}},{\"id\":\"3\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"status.raw\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\"}},{\"id\":\"4\",\"type\":\"terms\",\"schema\":\"segment\",\"params\":{\"field\":\"geoip.city_name.raw\",\"size\":5,\"order\":\"desc\",\"orderBy\":\"1\"}}],\"listeners\":{}}",
		//	"visState":              visStatetmp,
		"description":           " ",
		"title":                 domain + "-" + "uri" + id,
		"version":               1,
		"kibanaSavedObjectMeta": searchjson,
	}
	mJson, _ := json.Marshal(m)
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", url, contentReader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	return resp.StatusCode
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

func Nginxdashboard(domain string, topX int) (code int) {
	var vsList []string
	vsList = append(vsList, domain+"-base")

	for i := 1; i <= topX; i++ {
		vsList = append(vsList, domain+"-"+"uri"+strconv.Itoa(i))
	}

	for j, _ := range vsList {
		vsList[j] = strings.Replace(vsList[j], " ", "", -1)
	}

	//www.kaola.com-uri5
	//www.kaola.com-uri1

	url := NginxDash + domain + "-" + "dashboard"
	jsonstr := "{\"filter\":[{\"query\":{\"query_string\":{\"query\":\"*\",\"analyze_wildcard\":true}}}]}"
	searchjson := Json{SearchSourceJSON: jsonstr}

	fmt.Println("searchjson is", searchjson)
	//panelsjson := makePanelsjson(10, vsList)
	panelsjson_tem := makePanelsjson(topX, vsList)
	fmt.Println("======================")
	fmt.Println(panelsjson_tem)
	fmt.Println("======================")

	panelsjson := panelsjson_tem
	//panelsjson := fmt.Sprintf("[{\"col\":1,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"%s\",\"row\":2,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":2,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"%s\",\"row\":3,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":3,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"}]", "www.kaola.com-uri1", "www.kaola.com-uri1", "www.kaola.com-uri1", "www.kaola.com-uri1", "www.kaola.com-uri1", str)
	//panelsjson := fmt.Sprintf("[{\"col\":1,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"%s\",\"row\":2,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":2,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"%s\",\"row\":3,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":3,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"}]", vsList[0], vsList[1], vsList[2], vsList[3], vsList[4], vsList[5])
	m := map[string]interface{}{
		"hits":                  0,
		"panelsJSON":            panelsjson,
		"description":           " ",
		"title":                 domain + "-" + "dashboard",
		"version":               1,
		"kibanaSavedObjectMeta": searchjson,
	}
	fmt.Println("==")
	fmt.Println(m)
	fmt.Println("==")
	mJson, _ := json.Marshal(m)
	fmt.Println(string(mJson))
	contentReader := bytes.NewReader(mJson)
	req, _ := http.NewRequest("POST", url, contentReader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	return resp.StatusCode
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

}
