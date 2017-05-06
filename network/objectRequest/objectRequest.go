package objectRequest

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"os"
	"github.com/pankova/workStat/network/objectResponse"

)

// ObjectRequest - 
type ObjectRequest struct {
	Request 	string
}

func (obj *ObjectRequest) DoRequest(typename string, request string) string {
	obj.Request = request
	resultCanBePrinted := false
	response:= new(objectResponse.ObjectResponse)
	for resultCanBePrinted == false {
		response = obj.getResponseWithQuery(typename)
		if response.IsListOverThousand() {
			addParam := getParamFromConsole()
			obj.Request += "&q=" + addParam
		}else {
			resultCanBePrinted = true
		}
	}
	paramForNewRequest := response.GetSearchParameter()
	return paramForNewRequest
}

func (obj *ObjectRequest) getResponseWithQuery(typename string) *objectResponse.ObjectResponse {
	elems := new(objectResponse.ObjectResponse)
	elems.SetTypename(typename)

	resp, err := http.Get(obj.Request)
	if err != nil {
		log.Fatalf("Couldn't get: %v\n", err)
	}

	err = json.NewDecoder(resp.Body).Decode(elems)
	if err != nil {
		log.Fatalf("Error in parsing %s: %v\n", typename, err)
	}
	resp.Body.Close()

	fmt.Println("--- Распарсили " + typename + ":")

	if len(elems.Response.Items) == 0 {
		fmt.Print("--- Возможно, со времени последнего соединения прошло более суток и токен уже протух или " +
			"же в выбранном вами варианте нет объектов для поиска.")
		os.Exit(1)
	}
	return elems
}

func getParamFromConsole() string {
	var object string
	_, err := fmt.Scanf("%s", &object)
	if err != nil {
		log.Fatalf("Couldn't scanf: %v\n", err)
		return ""
	}
	return object
}
