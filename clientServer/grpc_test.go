package main

import (
	"context"
	"testing"

	pb "Danial-Movahed.github.io/clientServerGrpc"
)

func TestRequestHoliday(t *testing.T) {
	s := server{}

	// set up test cases
	type expectation struct {
		out *pb.Holiday
		err error
	}

	type Test struct {
		input    *pb.HolidayRequest
		expected expectation
	}

	tests := []Test{
		{
			input: &pb.HolidayRequest{
				Day:   "12",
				Month: "2",
				Year:  "2023",
			},
			expected: expectation{
				out: &pb.Holiday{
					Name:        "efee",
					Description: "loloooooo",
				},
				err: nil,
			},
		},
	}

	for _, test := range tests {
		resp, err := s.RequestHoliday(context.Background(), test.input)
		if err != nil {
			t.Errorf("Err -> \n\tWant: %q\n\tGot: %q\n", test.expected.err, err)
		}
		if resp != test.expected.out {
			t.Errorf("Body -> \n\twant %s, \n\tgot %s\n", test.expected.out, resp)

		}
	}
}
