package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// func compare(seats1, seats2 []string) bool {
// 	if len(seats1) != len(seats2) {
// 		return false
// 	}
// 	for i := range seats1 {
// 		if seats1[i] != seats2[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func getSeat(rows []string, x, y int) string {
// 	return string(rows[y][x])
// }

// func countOccupiedSeats(rows []string) (c int) {
// 	for y, r := range rows {
// 		for x := range r {
// 			//fmt.Println(getSeat(rows, x, y))
// 			if getSeat(rows, x, y) == occupied {
// 				c++
// 			}
// 		}
// 	}
// 	return c
// }

// func isOccupiedSeatInDirectionVisible(rows []string, x, y int, dir []int, depth int) bool {
// 	for {
// 		depth--
// 		x += dir[0]
// 		y += dir[1]
// 		if x < 0 || x >= len(rows[0]) || y < 0 || y >= len(rows) {
// 			break
// 		}
// 		if getSeat(rows, x, y) == occupied {
// 			return true
// 		}
// 		if getSeat(rows, x, y) == empty {
// 			break
// 		}
// 		if depth == 0 {
// 			break
// 		}
// 	}
// 	return false
// }

// func getOccupiedSeats(rows []string, x, y, depth int) int {
// 	occupieds := 0

// 	deltas := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
// 	for _, d := range deltas {
// 		if isOccupiedSeatInDirectionVisible(rows, x, y, d, depth) {
// 			occupieds++
// 		}
// 	}
// 	return occupieds
// }

// func getNewIteration(rows []string, tolerance, depth int) []string {
// 	var it []string
// 	for y, col := range rows {
// 		line := ""
// 		for x := range col {
// 			switch {
// 			case getSeat(rows, x, y) == empty && getOccupiedSeats(rows, x, y, depth) == 0:
// 				line = line + occupied
// 				break
// 			case getSeat(rows, x, y) == occupied && getOccupiedSeats(rows, x, y, depth) >= tolerance:
// 				line = line + empty
// 				break
// 			default:
// 				line = line + getSeat(rows, x, y)
// 				break
// 			}
// 		}
// 		it = append(it, line)
// 	}
// 	return it
// }

// func countSeatsForNeighbors(rows []string, tolerance, depth int) int {
// 	last := rows
// 	for {
// 		i := getNewIteration(last, tolerance, depth)
// 		if compare(last, i) {
// 			break
// 		}
// 		last = i

// 	}
// 	return countOccupiedSeats(last)
// }

// func countSeatsDirectNeighbors(rows []string) int {
// 	return countSeatsForNeighbors(rows, 4, 1)
// }

// func countSeatsVisibleNeighbors(rows []string) int {
// 	return countSeatsForNeighbors(rows, 5, -1)
// }

//---------

const (
	active = "#"
	//empty    = "L"
	inactive = "."
)

func getDirections() (directions [][]int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if !(i == 0 && j == 0 && k == 0) {
					directions = append(directions, []int{i, j, k})
				}
			}
		}
	}
	return directions
}

func getNeighbors(directions [][]int, x, y, z int) (neighbors [][]int) {
	for _, d := range directions {
		n := []int{x + d[0], y + d[1], z + d[2]}
		neighbors = append(neighbors, n)
	}
	return neighbors
}

func getNewActivity(cubes map[string]bool, current []int) bool {
	dir := getDirections()
	neighbors := getNeighbors(dir, current[0], current[1], current[2])
	var actives [][]int
	for _, n := range neighbors {
		if getByCoordinates(cubes, n) {
			actives = append(actives, n)
		}
	}
	l := len(actives)
	v := getByCoordinates(cubes, current)
	if v {
		if l == 2 || l == 3 {
			//fmt.Printf("%v %t : 1-true    %v\n ", current, v, actives)
			return true
		} else {
			//fmt.Printf("%v %t : 1-false    %v\n ", current, v, actives)
			return false
		}
	} else {
		if l == 3 {
			//fmt.Printf("%v %t : 2-true    %v\n ", current, v, actives)
			return true
		} else {
			//fmt.Printf("%v %t : 2-false    %v\n ", current, v, actives)
			return false
		}
	}
}

func coord2String(coordinate []int) string {
	s := fmt.Sprint(coordinate[0])
	for _, c := range coordinate[1:] {
		s = fmt.Sprintf("%s,%d", s, c)
	}
	return s
}

func string2Coord(cordinatesstrings string) (c []int) {
	cs := strings.Split(cordinatesstrings, ",")
	for _, s := range cs {
		i, _ := strconv.Atoi(string(s))
		c = append(c, i)
	}
	return c
}

func getByString(cubes map[string]bool, cs string) bool {
	pos, found := cubes[cs]
	if !found {
		return false
	}
	return pos
}

func getByCoordinates(cubes map[string]bool, coordinates []int) bool {
	return getByString(cubes, coord2String(coordinates))
}

func dim(n int) int {
	for i := 0; i < 1000; i++ {
		if int(math.Pow(float64(i), float64(3))) == n {
			return i
		}
	}
	return 0
}

func cubeString(cubes map[string]bool) {
	n := len(cubes)

	d := dim(n)

	zcorr := d/2 + 1
	//corr := int(math.Max(1, float64(d/2)))
	corr := d/2 + 1

	//fmt.Printf("n:%d  d:%d  c:%d", n, d, corr)
	for z := 0 - zcorr; z <= zcorr; z++ {
		fmt.Printf("z=%d\n", z)
		for y := 0 - corr; y <= corr; y++ {
			for x := 0 - corr; x <= corr; x++ {
				v := getByCoordinates(cubes, []int{x, y, z})
				if v {
					fmt.Print(active)
				} else {
					fmt.Print(inactive)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func getNewIteration(cubes map[string]bool) map[string]bool {
	iter := make(map[string]bool)
	n := len(cubes)
	d := dim(n) + 1
	fmt.Println(d)

	for z := -1; z <= d; z++ {
		for y := -1; y <= d; y++ {
			for x := -1; x <= d; x++ {
				coord := []int{x, y, z}

				a := getNewActivity(cubes, coord)
				c := coord2String([]int{x + 1, y + 1, z + 1})
				fmt.Printf("%v - %t - %v\n", coord, a, c)
				iter[c] = a
			}
		}
	}
	return iter
}

func cycle(cubes map[string]bool, n int) int {
	it := cubes
	for i := 0; i < n; i++ {
		it = getNewIteration(it)
	}

	count := 0
	for _, v := range it {
		if v {
			count++
		}
	}

	return count
}

func main() {

	var input string = "input.txt"
	cube := parsedata(readdata(input))
	fmt.Println(len(cube))
	fmt.Printf("Count seats occupied with direct neighborhood metric: %d\n", cycle(cube, 6))

	//fmt.Printf("Count seats occupied with visible seats metric:       %d\n", countSeatsVisibleNeighbors(seats))

}
