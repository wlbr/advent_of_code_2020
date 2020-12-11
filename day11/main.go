package main

import "fmt"

var input string = "input.txt"

const (
	occupied = "#"
	empty    = "L"
	noseat   = "."
)

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

func isOccupiedSeatInDirectionVisible(rows []string, x, y int, dir []int, depth int) bool {
	for {
		depth--
		x += dir[0]
		y += dir[1]
		if x < 0 || x >= len(rows[0]) || y < 0 || y >= len(rows) {
			break
		}
		if getSeat(rows, x, y) == occupied {
			return true
		}
		if getSeat(rows, x, y) == empty {
			break
		}
		if depth == 0 {
			break
		}
	}
	return false
}

func getOccupiedSeats(rows []string, x, y, depth int) int {
	occupieds := 0

	deltas := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, d := range deltas {
		if isOccupiedSeatInDirectionVisible(rows, x, y, d, depth) {
			occupieds++
		}
	}
	return occupieds
}

func getNewIteration(rows []string, tolerance, depth int) []string {
	var it []string
	for y, col := range rows {
		line := ""
		for x := range col {
			switch {
			case getSeat(rows, x, y) == empty && getOccupiedSeats(rows, x, y, depth) == 0:
				line = line + occupied
				break
			case getSeat(rows, x, y) == occupied && getOccupiedSeats(rows, x, y, depth) >= tolerance:
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

func countSeatsForNeighbors(rows []string, tolerance, depth int) int {
	last := rows
	for {
		i := getNewIteration(last, tolerance, depth)
		if compare(last, i) {
			break
		}
		last = i

	}
	return countOccupiedSeats(last)
}

func countSeatsDirectNeighbors(rows []string) int {
	return countSeatsForNeighbors(rows, 4, 1)
}

func countSeatsVisibleNeighbors(rows []string) int {
	return countSeatsForNeighbors(rows, 5, -1)
}

func main() {
	seats := readdata(input)

	fmt.Printf("Count seats occupied with direct neighborhood metric: %d\n", countSeatsDirectNeighbors(seats))

	fmt.Printf("Count seats occupied with visible seats metric:       %d\n", countSeatsVisibleNeighbors(seats))

}
