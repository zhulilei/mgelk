package libs

import (
	"fmt"
	_ "github.com/gorilla/mux"
	"net/http"
	//"strconv"
)

func Nginxtemp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")
	existItem, _ := parseRequest(w, r)
	fmt.Println(existItem)

	domain := existItem["domain"]
	index := existItem["index"]
	fromtime := existItem["from"]
	totime := existItem["to"]
	//topX := existItem["topX"]

	topX := 10

	if baseCode := NginxtempBase(index, domain, topX); baseCode != 201 {
		res := DataPack(baseCode, nil, "make-base-error")
		writeBody(res, w, http.StatusInternalServerError)
		return
	} else {
		domainUri := Geturi(index, domain, fromtime, totime, topX)
		fmt.Println("====")
		fmt.Println(domainUri)
		fmt.Println("====")
		for i, v := range domainUri[domain] {
			fmt.Println(v)
			if uricode := NginxtempUri(index, domain, v, i+1); uricode != 201 {
				res := DataPack(uricode, nil, "make-uri-error")
				writeBody(res, w, http.StatusInternalServerError)
				return
			}
		}
	}

	if dashCode := Nginxdashboard(domain, topX); dashCode != 201 {
		res := DataPack(dashCode, nil, "error")
		writeBody(res, w, http.StatusInternalServerError)
		return
	}
	res := DataPack(http.StatusCreated, nil, "success")
	writeBody(res, w, http.StatusCreated)
}
