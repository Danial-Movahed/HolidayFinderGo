package main

import (
	"testing"
)

type Test struct {
	input HolidayRequest
	want  Holiday
}

func TestGetHolidayRequest(t *testing.T) {
	// Create a request to pass to the handler
	var tests = []Test{
		{input: HolidayRequest{Day: "11", Month: "2", Year: "2002"}, want: Holiday{Name: "Revolution Day", Description: "Revolution Day is a national holiday in Iran"}},
		{input: HolidayRequest{Day: "28", Month: "7", Year: "2023"}, want: Holiday{Name: "Ashura", Description: "Ashura is a national holiday in Iran"}},
		{input: HolidayRequest{Day: "22", Month: "2", Year: "2222"}, want: Holiday{Name: "Nothing", Description: "No holidays on this date!"}},
		{input: HolidayRequest{Day: "35", Month: "1", Year: "1234"}, want: Holiday{Name: "Nothing", Description: "No holidays on this date!"}}}

	for _, test := range tests {
		got := get_holiday_request(test.input)
		if got != test.want {
			t.Errorf("handler returned wrong status code: want %s, got %s", test.want, got)
		}
	}
}
