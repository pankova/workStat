package objectResponse

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"github.com/pankova/workStat/entities/object"
	"github.com/pankova/workStat/network/responseError"
)

type ObjectResponse struct {
	Typename 	string
	Response 	object.Objects             		`json:"response"`
	Error    	*responseError.ResponseError	`json:"error,omitempty"`
}

func (objects *ObjectResponse) String() string {
	str := ""
	for i := 0; i < len(objects.Response.Items); i++ {
		object := objects.Response.Items[i]
		str += object.String() +"\n"
	}
	return str
}

func (obj *ObjectResponse) SetTypename(typename string) {
	obj.Typename = typename
}

func (elems *ObjectResponse) IsListOverThousand() bool {
	if elems.Response.Count > 1000 {
		fmt.Println("--- Результатов слишком много, чтобы не утомляться листанием длинного списка, введите," +
			"пожалуйста, первую букву или частичный запрос:")
		return true
	}
	if elems.Response.Count == 0 {
		fmt.Println("--- Введите, пожалуйста, другое уточнение")
		return true
	}
	return false
}


func (obj *ObjectResponse) getUserChoose() string {
	fmt.Printf("%v\n", obj)
	fmt.Println("--- Выберите из вариантов " + obj.Typename + "(можно ввести часть слова):")
	return getParamFromConsole()
}

func (obj *ObjectResponse) GetSearchParameter() string {
	param := obj.getUserChoose()

	for !(len(obj.Response.Items) == 1) {
		tempobj := obj.Response.Items
		*obj = obj.filterByUserParameter(param)
		if len(obj.Response.Items) > 1 {
			fmt.Println("--- Вашему запросу соответствует несколько результатов, введите, пожалуйста, более уникальную часть слова")
			param = obj.getUserChoose()
		}
		if len(obj.Response.Items) == 0 {
			fmt.Println("--- Результатов с такими данными нет. Выберите, пожалуйста, что-то другое:")
			obj.Response.Items = tempobj
			param = obj.getUserChoose()
		}
	}
	fmt.Println("--- Есть результат! Вы выбрали: " + obj.Response.Items[0].Name)
	return strconv.Itoa(obj.Response.Items[0].Id)
}

func (obj *ObjectResponse) filterByUserParameter(param string) ObjectResponse {
	filterObjects := new(ObjectResponse)
	items := filterObjects.Response.Items
	startSize := len(obj.Response.Items)
	for _, object := range obj.Response.Items {
		if object.HasParameterInName(param) {
			items = append(items, object)
		}
	}
	endSize := len(items)
	//ситуация, когда все оставшиеся варианты полностью включают подслово
	//например выбираем "Хабаровск", а список состоит из "Хабаровск" и "Хабаровск-43"
	//в таком случае ищем строку через полное совпадение (с точностью до регистра), а не по подстроке
	if startSize == endSize {
		for _, obj := range items {
			if strings.EqualFold(strings.ToLower(obj.Name), strings.ToLower(param)) {
				items = append(*new([]object.Object), obj)
			}
		}
		if len(items) > 1 {
			fmt.Println("--- Попробуйте, пожалуйста, ввести другое подслово.")
		}
	}
	filterObjects.Response.Items = items
	return *filterObjects
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


