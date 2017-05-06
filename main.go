package main

import (
	"encoding/json"
	"net/http"
	"log"
	"runtime"
	"github.com/pankova/workStat/entities/object/user"
	"github.com/pankova/workStat/entities/object/city"
	"github.com/pankova/workStat/config"
	"fmt"
	"os/exec"
	"github.com/pankova/workStat/network/objectRequest"
)

func open(url string) error {
    var cmd string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
    case "darwin":
        cmd = "open"
    default: // "linux", "freebsd", "openbsd", "netbsd"
        cmd = "xdg-open"
    }
    return exec.Command(cmd, url).Start()
}

func choose(users []user.User, filter func(user.User) bool) (filterUsers []user.User) {
	for _, user := range users {
		if filter(user) {
			filterUsers = append(filterUsers, user)
		}
	}
	return filterUsers
}

func HasCareer (user user.User) bool {
	if len(user.Work) > 0 {
		return true
	}
	return false
}

func main(){
	//authRequest := "https://oauth.vk.com/authorize?client_id=5812131&display=popup&redirect_uri=https://oauth.vk.com/blank.html&scope=friends&response_type=token&" + version
	//open(authRequest)

	setting := config.NewConfig()
	request := new(objectRequest.ObjectRequest)

	// ***** county ****** //
	countriesCode := setting.CountriesCode
	getCountries := setting.Common + "database.getCountries?" + setting.Version + setting.Tocken + "&code=" + countriesCode
	country := request.DoRequest("страны", getCountries)

	// ***** region ****** //
	getRegions := setting.Common + "database.getRegions?" + setting.Version + "&country_id=" + country + setting.Tocken + "&need_all=1" + "&count=1000"
	region := request.DoRequest("регионы", getRegions)

	// ***** city ****** //
	getCities := setting.Common + "database.getCities?" + setting.Version + "&country_id=" + country + "&region_id=" + region + setting.Tocken + "&count=1000"
	cityRequest := new(city.CityRequest)
	city := cityRequest.DoRequest("города", getCities)

	// ***** university ****** //
	getUniversities := setting.Common + "database.getUniversities?" + setting.Version + "&country_id=" + country + "&city_id=" + city + setting.Tocken + "&count=1000"
	university := request.DoRequest("университеты", getUniversities)

	// ***** faculty ****** //
	getFaculties := setting.Common + "database.getFaculties?" + setting.Version + "&university_id=" + university + setting.Tocken + "&count=1000"
	facultyRequest := objectRequest.ObjectRequest{Request: getFaculties}
	faculty := facultyRequest.DoRequest("факультеты", getFaculties)

	// Пример запроса, который требует апи вк. Обязателен только METHOD_NAME, оcтальное опционально
	// https://api.vk.com/method/METHOD_NAME?PARAMETERS&access_token=ACCESS_TOKEN&v=V
	getUsersBySearch := setting.Common + "users.search?" + setting.Version + "&university=" + university + "&university_faculty=" + faculty + setting.Tocken + "&count=1000" + "&fields=career"

	respu, err := http.Get(getUsersBySearch)
	if err != nil {
		log.Fatalf("Couldn't get: %v\n", err)
	}

	users := new(user.UserResponse)
	err = json.NewDecoder(respu.Body).Decode(users)
	if err != nil {
		log.Fatalf("Error in parsing users: %v\n", err)
	}
	usersWithCareer := choose(users.Response.Items, HasCareer)

	respu.Body.Close()
	fmt.Printf("Распарсили юзеров\n")

	if len(users.Response.Items) == 0 {
		fmt.Print("Возможно, со времени последнего соединения прошло более суток и токен уже протух.")
	}
	fmt.Printf("%v\n", users.Response.Count)

	users.Response = *users.Response.ParseGroupInfo()

	for _, user :=range usersWithCareer {
		fmt.Printf(user.String())
	}
}
