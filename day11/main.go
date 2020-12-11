package main

import "fmt"

var input string = "input.txt"

func Min(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e < m {
			m = e
		}
	}
	return m
}

func Max(v ...int) int {
	m := v[0]
	for _, e := range v {
		if e > m {
			m = e
		}
	}
	return m
}

const (
	occupied = "#"
	empty    = "L"
	noseat   = "."
)

func getSeat(rows []string, x, y int) string {
	return string(rows[y][x])
}

func countOccupiedSeats(rows []string) (c int) {
	for y, r := range rows {
		for x := range r {
			//fmt.Println(getSeat(rows, x, y))
			if getSeat(rows, x, y) == occupied {
				c++
			}
		}
	}
	return c
}

func isOccupiedSeatInDirectionVisible(rows []string, x, y int, dir []int) bool {
	for {
		x += dir[0]
		y += dir[1]
		if x < 0 || x >= len(rows[0]) || y < 0 || y >= len(rows) {
			break
		}
		if getSeat(rows, x, y) == occupied {
			return true
		}
		if getSeat(rows, x, y) == empty {
			return false
		}
	}
	return false
}

func getOccupiedSeats(rows []string, x, y int) int {
	occupieds := 0

	deltas := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range deltas {
		if isOccupiedSeatInDirectionVisible(rows, x, y, d) {
			occupieds++
		}
	}
	return occupieds
}

func getNewIterationTask2(rows []string) []string {
	var it []string
	for y, col := range rows {
		line := ""
		for x := range col {
			switch {
			case getSeat(rows, x, y) == empty && getOccupiedSeats(rows, x, y) == 0:
				line = line + occupied
				break
			case getSeat(rows, x, y) == occupied && getOccupiedSeats(rows, x, y) >= 5:
				line = line + empty
				break
			default:
				line = line + getSeat(rows, x, y)
				break
			}
		}
		it = append(it, line)
	}
	return it
}

func task2(rows []string) int {
	last := rows

	i := getNewIterationTask2(last)
	for {
		last = i
		i = getNewIterationTask2(last)

		//	fmt.Printf(" %+v\n %+v %t\n\n", last, i, compare(last, i))
		if compare(last, i) {
			break
		}
	}
	//fmt.Printf(" %+v\n", last)

	return countOccupiedSeats(last)
}

func compare(seats1, seats2 []string) bool {
	if len(seats1) != len(seats2) {
		return false
	}
	for i := range seats1 {
		if seats1[i] != seats2[i] {
			return false
		}
	}
	return true
}

func main() {
	seats := readdata(input)
	t1 := task1(seats)
	fmt.Println(t1)

	t2 := task2(seats)
	fmt.Println(t2)
}
