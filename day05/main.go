package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input string = "input.txt"

const COLS = 8
const ROWS = 128

func calcCode(row, col int) string {
	n := row*8 + col
	c := fmt.Sprintf("%10b\n", n)

	rr := strings.NewReplacer(" ", "F", "0", "F", "1", "B")
	code := rr.Replace(c[:6])

	cr := strings.NewReplacer("0", "L", "1", "R")
	code = code + cr.Replace(c[6:])

	return code
}

func calcSeatNo(code string) int {
	r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	code = r.Replace(code)
	c, _ := strconv.ParseInt(code, 2, 64)
	return int(c)
}

type seat struct {
	pass string
	id   int
	row  int
	col  int
}

func NewSeat(pass string) *seat {
	s := &seat{pass: pass}
	s.id = calcSeatNo(pass)
	s.row = s.id / COLS
	s.col = s.id % COLS
	return s
}

func (s *seat) String() string {
	return fmt.Sprintf("pass:'%s' id:'%d' row:'%d' col:'%d", s.pass, s.id, s.row, s.col)
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

func getMySeat(seats map[int]*seat) (myPossibleSeats []int) {
	for i := 1; i < COLS*ROWS-1; i++ {
		if seats[i] == nil && seats[i-1] != nil && seats[i+1] != nil {
			myPossibleSeats = append(myPossibleSeats, i)
		}
	}
	return myPossibleSeats
}

func getHighest(seats map[int]*seat) (max int) {
	for _, s := range seats {
		if s.id > max {
			max = s.id
		}
	}
	return max
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
	fmt.Printf("My seat is one of: %v\n", getMySeat(allSeats))

}
