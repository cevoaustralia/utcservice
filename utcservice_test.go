package main

import (
	"testing"
)

var dateTests = []struct {
	input   string
	tz      string
	isError bool
	output  string
}{
	{"1540437516", "", false, "2018-10-25 03:18:36 +0000 UTC"},
	{"1540437516", "Australia/Melbourne", false, "2018-10-25 14:18:36 +1100 AEDT"},
	{"bleem", "", true, "1970-01-01 00:00:00 +0000 UTC"},
	{"bleem", "Australia/Melbourne", true, "1970-01-01 10:00:00 +1000 AEST"},
}

func TestParseDate(t *testing.T) {
	for _, tt := range dateTests {
		t.Run(tt.input, func(t *testing.T) {
			out, err := parseDate(tt.input, tt.tz)
			if err != nil && !tt.isError {
				t.Errorf("got error, did not want")
			}
			if out.String() != tt.output {
				t.Errorf("got %s, wanted %s", out.String(), tt.output)
			}
		})
	}
}
