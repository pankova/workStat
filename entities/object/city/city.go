package city

import (
	"github.com/pankova/workStat/entities/object"
	"github.com/pankova/workStat/network/objectRequest"
	"github.com/pankova/workStat/network/objectResponse"


)
// City - расширенный базовый объект Object с дополнительной информацией (https://vk.com/dev/database.getCities)
type City struct {
	object.Object
	Area		string		`json:"area"`
	Region		string		`json:"region"`
	// Important (=1) отмечены ключевые города для текущего пользователя
	Important	int		`json:"important"`
}

type CityRequest struct {
	objectRequest.ObjectRequest
	objectResponse.ObjectResponse
}

func (city *City) String() string {
	str := ""
	str += city.Object.String() +
		" " + city.Region +
		" " + city.Area
	return str
}