package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type datetime struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}
type date struct {
	Iso      string   `json:"iso"`
	Datetime datetime `json:"datetime"`
}
type holiday struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Date        date     `json:"date"`
	Type        []string `json:"type"`
}
type meta struct {
	Code int `json:"code"`
}
type response struct {
	Holidays []holiday `json:"holidays"`
}
type Res struct {
	Meta     meta     `json:"meta"`
	Response response `json:"response"`
}
type HolidayRequest struct {
	Day   string
	Month string
	Year  string
}
type Holiday struct {
	Name        string
	Description string
}

func get_holiday_request(req HolidayRequest) (holi Holiday) {
	requestURL := "https://calendarific.com/api/v2/holidays?api_key=173842ca0f6299a6ec40c835b958e49f1d63548f&country=IR&year=" + req.Year
	res, err := http.Get(requestURL) // a = requests.get()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var holidays Res

	errr := json.Unmarshal(resBytes, &holidays)
	if errr != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	number_of_holidays := len(holidays.Response.Holidays)
	//a := int(holidays.Response.Holidays[10].Date.Datetime.Year)
	//fmt.Println(a)
	year, errr := strconv.Atoi(req.Year)
	if errr != nil {
		fmt.Println("Error during conversion")
		return
	}
	month, errrr := strconv.Atoi(req.Month)
	if errrr != nil {
		fmt.Println("Error during conversion")
		return
	}
	day, errrrr := strconv.Atoi(req.Day)
	if errrrr != nil {
		fmt.Println("Error during conversion")
		return
	}
	//fmt.Println(number_of_holidays)
	for i := 0; i < number_of_holidays; i++ {
		fmt.Println(holidays.Response.Holidays[i].Date.Datetime)
		if int(holidays.Response.Holidays[i].Date.Datetime.Year) == year && int(holidays.Response.Holidays[i].Date.Datetime.Month) == month && int(holidays.Response.Holidays[i].Date.Datetime.Day) == day {

			holi.Name = holidays.Response.Holidays[i].Name
			holi.Description = holidays.Response.Holidays[i].Description
			fmt.Println(i)
			return holi
		}

	}
	//fmt.Println(holidays.Response.Holidays)
	holi.Name = "Not Holiday"
	holi.Description = ""
	return holi
}

func main() {
	fmt.Println(get_holiday_request(HolidayRequest{"27", "12", "2022"}))

}
