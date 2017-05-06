package user

import (
	"strconv"
	"github.com/pankova/workStat/entities/object/city"
	"github.com/pankova/workStat/config"
	"github.com/pankova/workStat/entities/group"
	"net/http"
	"log"
	"encoding/json"
)

type User struct {
	Id		int32			`json:"id"`
	Name		string			`json:"first_name"`
	Lastname	string			`json:"last_name"`
	Work		[]Career		`json:"career"`
	City 		city.City 		`json:"city"`
}

type Career struct {
	GroupId  	int 			`json:"group_id"`
	Group		group.GroupResponse
	Company  	string			`json:"company"`
	Position 	string			`json:"position"`
	City     	string			`json:"city_name"`
}

type Users struct {
	Items		[]User			`json:"items"`
	Count		int32 			`json:"count"`
}

type UserResponse struct {
	Response	Users			`json:"response"`
	Error		*ResponseError	`json:"error,omitempty"`
}

type ResponseError struct {
	Code 		int				`json:"error_code"`
	Message		string			`json:"error_msg"`
}

func (users *UserResponse) String() string {
	str := ""
	for i := 0; i < len(users.Response.Items); i++ {
		user := users.Response.Items[i]
		str += user.String()
		str += "\n"
	}
	return str
}

func (user *User) String() string {
	str := ""

	city := user.City.String()
	if len(city) > 0 {
		str += city
	}

	if len(user.Work) > 0 {
		cityWork := user.Work[0].City
		company := user.Work[0].Company
		group := user.Work[0].Group.String()
		position := user.Work[0].Position

		if len(cityWork) > 0 {
			str += "Город: " + cityWork + " "
		}
		if len(company) > 0 {
			str += "Компания: " + company + " "
		}
		if len(group) > 0 {
			str += group + " "
		}
		if len(position) > 0 {
			str += "Должность: " + position + " "
		}
	}
	str += "\n"
	return str
}

func (users *Users) ParseGroupInfo() *Users {
	setting := config.NewConfig()
	for _,user := range users.Items {
		if len(user.Work) > 0 {
			if user.Work[0].GroupId > 0 {
				getGroupInfo := setting.Common + "groups.getById?" + setting.Version + setting.Tocken +
					"&group_ids=" + strconv.Itoa(user.Work[0].GroupId) + "&fields=description"

				resp, err := http.Get(getGroupInfo)
				if err != nil {
					log.Fatalf("Couldn't get: %v\n", err)
				}

				err = json.NewDecoder(resp.Body).Decode(&user.Work[0].Group)
				if err != nil {
					log.Fatalf("Error in parsing: %v\n", err)
				}
				resp.Body.Close()

			}
		}
	}
	return users
}
