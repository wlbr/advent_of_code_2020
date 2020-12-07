package main

import (
	"testing"
)

type testdata struct {
	fname         string
	input         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", "shiny gold", 4, 32}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		rules := getRules(readAllRules(test.fname))
		c := len(whoCanCarryThis(test.input, rules))
		if c != test.expectedtask1 {
			t.Fatalf("Test1 '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		rules := getRules(readAllRules(test.fname))
		c := howManyBags(test.input, rules)
		if c != test.expectedtask2 {
			t.Fatalf("Test2 '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}
