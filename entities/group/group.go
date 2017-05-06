package group

import (
	"strings"
	//"github.com/pankova/workStat/entities/object/city"
	"github.com/pankova/workStat/network/responseError"
	)

// Group - объект, описывающий сообщество ВК (https://vk.com/dev/groups)
type Group struct {
	Name			string 		`json:"name"`
	Description 	string 		`json:"description"`
}

func (group *Group) String() string {
	if len(group.Name) == 0 {
		return ""
	}
	str := ""
	str += "Группа ВКонтакте: " + group.Name

	description := group.Description
	if len(description) > 0 {
		if len(description) > 50 {
			description = description[:50]
			description = description + "..."
		}
		description = strings.Replace(description, "\n", "", -1)
		str += " Описание: " + description
	}
	return str
}

type GroupResponse struct {
	Response 	[]Group        					`json:"response"`
	Error    	*responseError.ResponseError	`json:"error,omitempty"`
}

func (group *GroupResponse) String() string {
	str := ""
	if len(group.Response) > 0{
		str += group.Response[0].String()
	}
	return str
}
