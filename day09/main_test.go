package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", 127, 62}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := lookForFirstNotSum(m, 5)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		n := lookForFirstNotSum(m, 5)
		c := lookForSumOfContiguousSum(m, n)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}
