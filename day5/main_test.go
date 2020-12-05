package main

import (
	"testing"
)

var testset []*seat = []*seat{&seat{pass: "FBFBBFFRLR", id: 357, row: 44, col: 5},
	&seat{pass: "BFFFBBFRRR", id: 567, row: 70, col: 7},
	&seat{pass: "FFFBBBFRRR", id: 119, row: 14, col: 7},
	&seat{pass: "BBFFBBFRLL", id: 820, row: 102, col: 4}}

func TestOne(t *testing.T) {
	for _, test := range testset {
		s := NewSeat(test.pass)
		if s.id != test.id || s.row != test.row || s.col != test.col {
			t.Fatalf("`Test %s failed.  %s  (Wanted: %s)'", test.pass, s, test)
		}
	}
}
