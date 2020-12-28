package main

import (
	"testing"
)

// type testdata struct {
// 	fname         string
// 	expectedtask1 []string
// }

//var testset1 *testdata = &testdata{"example.txt", []string{"iteration1-1.txt", "iteration1-2.txt", "iteration1-3.txt", "iteration1-4.txt", "iteration1-5.txt"}}
// var testset2 *testdata = &testdata{"example.txt", []string{"iteration2-1.txt", "iteration2-2.txt", "iteration2-3.txt", "iteration2-4.txt", "iteration2-5.txt", "iteration2-6.txt"}}

type testdata struct {
	fname         string
	expectedtask1 int
}

var testset1 *testdata = &testdata{"example.txt", 112}

func TestTaskOne(t *testing.T) {
	cube := parsedata(readdata(testset1.fname))
	c := cycle(cube, 1)
	exait := testset1.expectedtask1
	if c != exait {
		t.Fatalf("Test '%s' failed.  Got: '%d' \nWanted: '%d'", testset1.fname, c, exait)
	}
}

// func TestCountOccupiedSeats(t *testing.T) {
// 	fname := "iteration1-5.txt"
// 	expected := 37

// 	seats := readdata(fname)
// 	c := countOccupiedSeats(seats)
// 	if c != expected {
// 		t.Fatalf("TestCountOccupiedSeats failed. In %s:\n   Got: '%d' \nWanted: '%d'", fname, c, expected)
// 	}
// }

// func inspect(old, it, expected []string) {
// 	for i := range old {
// 		fmt.Printf("%s  %s  %s\n", old[i], it[i], expected[i])
// 	}
// }

// func TestTaskTwo(t *testing.T) {
// 	new := readdata(testset1.fname)

// 	for _, iteration := range testset2.expectedtask1 {
// 		old := new
// 		new = getNewIteration(new, 5, -1)
// 		expected := readdata(iteration)
// 		c := compare(new, expected)
// 		if !c {
// 			inspect(old, new, expected)
// 			t.Fatalf("Test '%s' failed. In %s:\n   Got: '%s' \nWanted: '%s'", testset2.fname, iteration, new, expected)
// 		}
// 	}
// }
