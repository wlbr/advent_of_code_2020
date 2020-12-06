package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", 11, 6}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		groups := readAllGroups(test.fname)
		c := sumOverAllGroups(groups, countAnswersPerGroupAnyone)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		groups := readAllGroups(test.fname)
		c := sumOverAllGroups(groups, countAnswersPerGroupEveryone)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}
