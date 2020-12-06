package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var input string = "input.txt"

const COLS = 8
const ROWS = 128

type seat struct {
	pass string
	id   int
	row  int
	col  int
}

func NewSeat(pass string) *seat {
	s := &seat{pass: pass}
	s.row, s.col = getPos(pass)
	s.id = calcID(s.row, s.col)
	return s
}

func (t *seat) String() string {
	return fmt.Sprintf("pass:'%s' id:'%d' row:'%d' col:'%d", t.pass, t.id, t.row, t.col)
}

func calcID(row, col int) int {
	return row*8 + col
}

func binsearch(code string, rows bool) int {
	max := int(math.Pow(2, float64(len(code))))
	lr := 0
	mr := max - 1
	discriminator := "f"
	if !rows {
		discriminator = "l"
	}

	for _, c := range code {
		k := strings.ToLower(string(c))
		if k == discriminator {
			mr = mr - ((mr - lr) / 2) - 1
		} else {
			lr = lr + ((mr - lr) / 2) + 1
		}
	}
	if rows {
		return lr
	}
	return mr
}

// magic happens in here
func getPos(p string) (row, col int) {
	row = binsearch(p[:7], true)
	col = binsearch(p[7:], false)

	return row, col

}

func getHighest(seats map[int]*seat) (max int) {
	for _, s := range seats {
		if s.id > max {
			max = s.id
		}
	}
	return max
}

func readInput(input string) (passids []string) {
	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		passids = append(passids, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return passids

}

func makeAllSeats(rows, cols int) [][]int {
	allseats := make([][]int, rows)
	for i := range allseats {
		allseats[i] = make([]int, cols)
	}
	return allseats
}

func getEmptyseats(seats map[int]*seat) (empties []int) {
	allseats := makeAllSeats(ROWS, COLS)
	for _, s := range seats {
		allseats[s.row][s.col] = 1
	}
	for r, row := range allseats {
		for c, s := range row {
			if s == 0 {
				empties = append(empties, r*COLS+c)
			}
		}
	}
	return empties
}

func getMySeat(seats map[int]*seat) (myPossibleSeats []int) {
	empties := getEmptyseats(seats)
	for i := 1; i < len(empties)-1; i++ {
		c := empties[i]
		if empties[i-1] != c-1 && empties[i+1] != c+1 {
			myPossibleSeats = append(myPossibleSeats, c)
		}
	}
	return myPossibleSeats
}

func main() {
	in := readInput(input)
	allSeats := make(map[int]*seat)
	for _, i := range in {
		s := NewSeat(i)
		allSeats[s.id] = s
	}
	max := getHighest(allSeats)
	fmt.Printf("The seat with the highest id is seat no: '%d'\n", max)
	fmt.Printf("My seat is %v.\n", getMySeat(allSeats))

}
