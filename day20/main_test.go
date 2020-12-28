package main

import (
	"testing"
)

type testdata struct {
	fname1        string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 20899048083289, 0},
	{"example2.txt", 20899048083289, 0},
	{"example3.txt", 20899048083289, 0}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		tiles := ReadData(test.fname1)
		c := t1(tiles)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname1, c, test.expectedtask1)
		}
	}
}

// func TestTaskTwo(t *testing.T) {
// 	for _, test := range testset {
// 		rules, myticket, nearbies := ParseData(ReadData(test.fname2))
// 		c := ProductOfDepartureField(rules, myticket, nearbies)
// 		if c != test.expectedtask2 {
// 			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname2, c, test.expectedtask2)
// 		}
// 	}
// }
