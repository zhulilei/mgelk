package libs

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func writeBody(item interface{}, w http.ResponseWriter, status int) {
	if body, err := json.Marshal(item); err == nil {
		w.WriteHeader(status)
		io.WriteString(w, string(body))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func parseRequest(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	var item map[string]string
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return item, err
	}
	return item, nil
}

func DataPack(code int, data interface{}, message string) (result Result) {
	result.Data = data
	result.Code = code
	result.Message = message
	return
}

func DealTime(timestring string) string {
	t, _ := time.Parse("2006/01/02-15:04", timestring) //time.time
	//fmt.Println(t, err)
	d, _ := time.ParseDuration("-8h")
	timestamp := t.Add(d)
	//fmt.Println(timestamp, err)
	return timestamp.Format("2006/01/02-15:04") //string
}

/*
panelsjson := fmt.Sprintf("[{\"col\":1,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"%s\",\"row\":2,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":2,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":1,\"id\":\"%s\",\"row\":3,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"},{\"col\":7,\"id\":\"%s\",\"row\":3,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"}]", vsList[0], vsList[1], vsList[2], vsList[3], vsList[4], vsList[5])
*/

func makePanelsjson(number int, vsList []string) (list string) {
	var jsonList []string
	list = ""
	//a := "{\\\"c ol\\\":1,\\\"id\\\":\\\"%s\",\\\"row\\\":1,\\\"size_x\\\":6,\\\"size_y\\\":4,\\\"type\\\":\\\"visualization\\\"}"

	for v := 1; v <= number; v++ {
		//fmt.Println("=--")
		//panelsjson := fmt.Sprintf("{\\\"col\\\":1,\\\"id\\\":\\\"%s\\\",\\\"row\\\":%d,\\\"size_x\\\":6,\\\"size_y\\\":4,\\\"type\\\":\\\"visualization\\\"}", vsList[v-1], v)
		panelsjson := fmt.Sprintf("{\"col\":1,\"id\":\"%s\",\"row\":%d,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"}", vsList[v-1], v)
		//fmt.Println(panelsjson)
		jsonList = append(jsonList, panelsjson)
		//fmt.Println("=--")
	}

	//panelsjson := fmt.Sprintf("{\\\"col\\\":%d,\\\"id\\\":\\\"%s\\\",\\\"row\\\":1,\\\"size_x\\\":6,\\\"size_y\\\":4,\\\"type\\\":\\\"visualization\\\"}", 1, "vlist1")

	//fmt.Println(jsonList)
	for _, v := range jsonList[:len(jsonList)-1] {
		list = list + v + ","
	}
	list = list + jsonList[len(jsonList)-1]
	fmt.Println("[" + list + "]")
	alist := "[" + list + "]"
	return alist

	//panelsjson := fmt.Sprintf("[{\"col\":1,\"id\":\"%s\",\"row\":1,\"size_x\":6,\"size_y\":4,\"type\":\"visualization\"}]", "hahah")

}
