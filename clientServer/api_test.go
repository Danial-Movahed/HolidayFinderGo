package main

import (
	"errors"
	"testing"
)

type expectation struct {
	out Holiday
	err error
}

type Test struct {
	input HolidayRequest
	want  expectation
}

func TestGetHolidayRequest(t *testing.T) {
	var tests = []Test{
		{
			input: HolidayRequest{
				Day:   "11",
				Month: "2",
				Year:  "2002",
			},
			want: expectation{
				out: Holiday{
					Name:        "Revolution Day",
					Description: "Revolution Day is a national holiday in Iran",
				},
				err: nil,
			},
		},
		{
			input: HolidayRequest{
				Day:   "28",
				Month: "7",
				Year:  "2023",
			},
			want: expectation{
				out: Holiday{
					Name:        "Ashura",
					Description: "Ashura is a national holiday in Iran",
				},
				err: nil,
			},
		},
		{
			input: HolidayRequest{
				Day:   "22",
				Month: "2",
				Year:  "2222",
			},
			want: expectation{
				out: Holiday{
					Name:        "Nothing",
					Description: "No holidays on this date!",
				},
				err: nil,
			},
		},
		{
			input: HolidayRequest{
				Day:   "21",
				Month: "30",
				Year:  "4321",
			},
			want: expectation{
				out: Holiday{
					Name:        "Nothing",
					Description: "No holidays on this date!",
				},
				err: errors.New("error"),
			},
		},
		{
			input: HolidayRequest{
				Day:   "35",
				Month: "1",
				Year:  "1234",
			},
			want: expectation{
				out: Holiday{
					Name:        "Nothing",
					Description: "No holidays on this date!",
				},
				err: errors.New("error"),
			},
		},
	}

	for _, test := range tests {
		got, err := get_holiday_request(test.input)
		if err != nil {
			if test.want.err.Error() != err.Error() {
				t.Errorf("Err -> \n\tWant: %q\n\tGot: %q\n", test.want.err, err)
			}
		} else {
			if got != test.want.out {
				t.Errorf("Body -> \n\twant %s, \n\tgot %s\n", test.want.out, got)
			}
		}
	}
}
