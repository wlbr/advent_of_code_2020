package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 []int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example.txt", []int{17, 8, 25}, 286},
	{"example2.txt", []int{-27, -14, 41}, 434}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		cmds := parsedata(readdata(test.fname))
		_, _, r := travel(cmds)
		if r != test.expectedtask1[2] {
			t.Fatalf("Test '%s' failed.  Got: '%d' Wanted: '%d'", test.fname, r, test.expectedtask1[2])
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		cmds := parsedata(readdata(test.fname))
		_, _, r := waypointing(cmds)
		if r != test.expectedtask2 {
			t.Fatalf("Test '%s' failed.  Got: '%d' Wanted: '%d'", test.fname, r, test.expectedtask2)
		}
	}
}
