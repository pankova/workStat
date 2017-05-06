package object

import (
	"strings"
)

// Object - базовый объект для сущностей вроде страны, региона, университета (то есть тех, у кого нам интересны только айдишник и имя).
// Если требуется более продвинутый объект, содержащий больше информации,
// то он включает в себя Object в качестве поля (на данный момент это City и User).
type Object struct {
	Id 			int         `json:"id"`
	Name    	string      `json:"title"`
}

type Objects struct {
	Items       []Object    `json:"items"`
	Count 		int32 		`json:"count"`
}

func (object *Object) String() string {
	str := ""
	str += //strconv.Itoa(int(object.Id)) + " " +
		object.Name
	return str
}

func (obj *Object) HasParameterInName(param string) bool {
	if strings.Contains(strings.ToLower(obj.Name), strings.ToLower(param)) {
		return true
	}
	//fmt.Printf("\nЯ не нашел " + param + " в " + obj.Name + "\n")
	return false
}
