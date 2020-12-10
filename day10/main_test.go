package main

import (
	"testing"
)

type testdata struct {
	fname         string
	expectedtask1 int
	expectedtask2 int
}

var testset []*testdata = []*testdata{{"example1.txt", 7 * 5, 8}, {"example2.txt", 22 * 10, 19208}}

func TestTaskOne(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := getProductOfCountOfDistances(m)
		if c != test.expectedtask1 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask1)
		}
	}
}

func TestTaskTwo(t *testing.T) {
	for _, test := range testset {
		m := readdata(test.fname)
		c := countCombinations(m)
		if c != test.expectedtask2 {
			t.Fatalf("Test '%s' failed. Got '%d' -  Wanted: '%d'", test.fname, c, test.expectedtask2)
		}
	}
}

type testcandidate struct {
	in  int
	out []int
}

var testcandidates []*testcandidate = []*testcandidate{{0, []int{1}}, {1, []int{4}}, {4, []int{5, 6, 7}}, {7, []int{10}}, {10, []int{11, 12}}, {16, []int{19}}, {19, []int{}}}

func allEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func TestGetCandidates(t *testing.T) {
	nums := readdata("example1.txt")
	for _, test := range testcandidates {
		c := getCandidates(nums, test.in)
		if !allEqual(c, test.out) {
			t.Fatalf("Test getCandidates('%d') failed. Got '%v' -  Wanted: '%v'", test.in, c, test.out)
		}
	}
}
