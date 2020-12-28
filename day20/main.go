package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func mult(nums ...int) int {
	prod := 1
	for _, n := range nums {
		prod *= n
	}
	return prod
}

type Tile struct {
	dim    []int
	id     int
	bfield [][]bool
	field  []string
	//original  []string
	sides []int
	//permstate int
}

func (t *Tile) String() string {
	s := fmt.Sprintf("Tile %d:", t.id)
	for _, l := range t.bfield {
		s += fmt.Sprintf("\n")
		for _, f := range l {
			if f {
				s += "#"
			} else {
				s += "."
			}
		}
	}
	return s
}

// func (t *Tile) North() int { return t.sides[0] }
// func (t *Tile) East() int  { return t.sides[1] }
// func (t *Tile) South() int { return t.sides[2] }
// func (t *Tile) West() int  { return t.sides[3] }

// func (t *Tile) FitsToTheTopOf(t2 *Tile) bool    { return t.South() == t2.North() }
// func (t *Tile) FitsToTheRightOf(t2 *Tile) bool  { return t.West() == t2.East() }
// func (t *Tile) FitsToTheBottomOf(t2 *Tile) bool { return t.North() == t2.South() }
// func (t *Tile) FitsToTheLeftOf(t2 *Tile) bool   { return t.East() == t2.West() }

// func (t *Tile) FitsInto(newtiles []*Tile, dim int) bool {
// 	if len(newtiles) == 0 {
// 		return true
// 	}
// 	if len(newtiles) < dim {
// 		return t.FitsToTheRightOf(newtiles[len(newtiles)-1])
// 	}
// 	if len(newtiles)%dim == 0 {
// 		return t.FitsToTheBottomOf(newtiles[len(newtiles)-dim])
// 	}

// 	cl := t.FitsToTheRightOf(newtiles[len(newtiles)-1])
// 	cr := t.FitsToTheBottomOf(newtiles[len(newtiles)-dim])

// 	c := cl && cr
// 	return c
// }

var titlerex = regexp.MustCompile(`^.* +(\d+):$`)

func getId(tiletitle string) (id int, e error) {
	matches := titlerex.FindStringSubmatch(tiletitle)
	if len(matches) >= 2 {
		if i, err := strconv.Atoi(matches[1]); err == nil {
			id = i
		} else {
			e = fmt.Errorf("tile id is not a numer: '%s'", tiletitle)
		}
	} else {
		e = fmt.Errorf("tile not in standard format: '%s'", tiletitle)
	}
	return id, e
}

func getField(field []string) [][]bool {
	var bfield [][]bool
	for _, line := range field {
		var bline []bool
		for _, c := range line {
			if c == '.' {
				bline = append(bline, false)
			} else {
				bline = append(bline, true)
			}
		}
		bfield = append(bfield, bline)
	}
	return bfield
}
func side2bin(side []bool) int {
	n := ""
	for _, s := range side {
		if s {
			n += "1"
		} else {
			n += "0"
		}
	}
	var i int64
	var err error
	if i, err = strconv.ParseInt(n, 2, 64); err != nil {
		i = -1
		log.Print(err)
	}
	return int(i)
}

func getSides(field [][]bool) []int {
	var sides []int
	sides = append(sides, side2bin(field[0]))

	var right []bool
	for _, line := range field {
		right = append(right, line[len(line)-1])
	}
	sides = append(sides, side2bin(right))

	sides = append(sides, side2bin(field[len(field)-1]))
	var left []bool
	for _, line := range field {
		left = append(left, line[0])
	}
	sides = append(sides, side2bin(left))

	return sides
}

func NewTile(id string, field []string) *Tile {
	t := &Tile{field: field}
	t.id, _ = getId(id)
	//t.field = append(t.field, t.original...)
	t.bfield = getField(field)
	t.dim = []int{len(field), len(field[0])}
	t.sides = getSides(t.bfield)
	return t
}

// func (t *Tile) Reset() {
// 	t.field = []string{}
// 	t.field = append(t.field, t.original...)
// 	t.permstate = 0
// 	t.bfield = getField(t.field)
// 	t.dim = []int{len(t.field), len(t.field[0])}
// 	t.sides = getSides(t.bfield)
// }

// func (t *Tile) Rotate() {
// 	var res [][]byte
// 	for _, line := range t.field {
// 		res = append(res, make([]byte, len(line)))
// 	}
// 	n := len(t.field)
// 	x := n / 2
// 	y := n - 1
// 	for i := 0; i < x; i++ {
// 		for j := i; j < y-i; j++ {
// 			k := t.field[i][j]
// 			res[i][j] = t.field[y-j][i]
// 			res[y-j][i] = t.field[y-i][y-j]
// 			res[y-i][y-j] = t.field[j][y-i]
// 			res[j][y-i] = k
// 		}
// 	}
// 	t.field = nil
// 	for _, line := range res {
// 		t.field = append(t.field, string(line))
// 	}
// 	t.bfield = getField(t.field)
// 	t.sides = getSides(t.bfield)
// }

// func (t *Tile) Flipv() {
// 	var res []string

// 	for i := len(t.field) - 1; i >= 0; i-- {
// 		res = append(res, t.field[i])
// 	}
// 	t.field = res
// 	t.bfield = getField(t.field)
// 	t.sides = getSides(t.bfield)
// }

// func reverse(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// func (t *Tile) Fliph() {
// 	var res []string
// 	for _, f := range t.field {
// 		res = append(res, reverse(f))
// 	}
// 	t.field = res
// 	t.bfield = getField(t.field)
// 	t.sides = getSides(t.bfield)
// }

// func PrintSquareIDs(tiles []*Tile) {
// 	d := int(math.Sqrt(float64(len(tiles))))
// 	for i := 0; i < d; i++ {
// 		for j := 0; j < d; j++ {
// 			fmt.Print(tiles[d*i+j].id, " ")
// 		}
// 		fmt.Println()
// 	}
// }

// func PrintSquareFields(tiles []*Tile) {
// 	d := int(math.Sqrt(float64(len(tiles))))
// 	for i := 0; i < d; i++ {
// 		for j := 0; j < len(tiles[0].field)-1; j++ {
// 			for k := 0; k < d; k++ {
// 				fmt.Print(tiles[d*i+k].field[j], " ")
// 			}
// 			fmt.Println()
// 		}
// 		fmt.Println()
// 	}
// }

// func Reset(tiles []*Tile) {
// 	for _, t := range tiles {
// 		t.Reset()
// 	}
// }

// func GetCorners(tiles []*Tile) []int {
// 	if len(tiles) == 0 {
// 		return []int{0, 0, 0, 0}
// 	}
// 	d := int(math.Sqrt(float64(len(tiles))))
// 	return []int{tiles[0].id, tiles[d-1].id, tiles[len(tiles)-d].id, tiles[len(tiles)-1].id}
// }

// const MAXPERMS = 8

// func (t *Tile) Permutate() {
// 	switch t.permstate {
// 	case 0, 1, 2, 4, 5, 6:
// 		t.Rotate()
// 		t.permstate++
// 		break
// 	// case 4, 9:
// 	// 	t.Fliph()
// 	// 	t.permstate++
// 	// 	break
// 	case 3, 7:
// 		t.Flipv()
// 		t.permstate++
// 		break
// 	default:
// 		t.permstate = 0
// 		break
// 	}
// }

// func (t *Tile) tryPermutations(nt []*Tile, d int) bool {
// 	if t.FitsInto(nt, d) {
// 		return true
// 	} else {
// 		for i := 0; i < MAXPERMS; i++ {
// 			t.Permutate()
// 			if t.FitsInto(nt, d) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func rec(remaining, already []*Tile, dim int) (found []*Tile) {
// 	if len(remaining) == 0 {
// 		return already
// 	}
// 	for i := 0; i < len(remaining); i++ {
// 		if remaining[i].tryPermutations(already, dim) {
// 			newalready := append(already, remaining[i])
// 			newremaining := append(remaining[:i], remaining[i+1:]...)
// 			found = append(found, rec(newremaining, newalready, dim)...)
// 		}
// 	}
// 	return found
// }

// func show(t *Tile, rest []*Tile) {
// 	fmt.Print("[ ", t.id, "-")
// 	for _, v := range rest {
// 		fmt.Print(" ", v.id)

// 	}
// 	fmt.Println(" ]")
// }

// func inspect(t *Tile, r1, r2 []*Tile) {
// 	fmt.Printf("t0: %d   r1 ", t.id)
// 	for _, v := range r1 {
// 		fmt.Printf(" %d", v.id)
// 	}
// 	fmt.Printf("  r2: ")
// 	for _, v := range r2 {
// 		fmt.Printf(" %d", v.id)
// 	}
// 	fmt.Println()
// }

func find(n int, nums []int) bool {
	for _, i := range nums {
		if n == i {
			return true
		}
	}
	return false
}

func t1(tiles []*Tile) int {
	//d := int(math.Sqrt(float64(len(tiles))))

	occurences := make(map[int]int)
	for _, t := range tiles {
		for _, s := range t.sides {
			occurences[s] = occurences[s] + 1
		}
	}
	var uniqs []int
	for u, oc := range occurences {
		if oc == 1 {
			uniqs = append(uniqs, u)
		}
	}

	occurences = make(map[int]int)
	for _, t := range tiles {
		for _, u := range uniqs {
			if find(u, t.sides) {
				occurences[t.id] = occurences[t.id] + 1
			}
		}
	}

	var corners []int
	for tid, oc := range occurences {
		if oc == 2 {
			corners = append(corners, tid)
		}

	}

	fmt.Println("-", corners, "-")

	return 0
}

func main() {
	var input string = "input.txt"
	// input = "example2.txt"
	tiles := ReadData(input)

	fmt.Println("t1 ", t1(tiles))
	// fmt.Printf("ticket scanning error rate: %d\n", tickedScanningErrorRate(rules, myticket, nearbies))
	// fmt.Printf("Prouct of the six departure numbers: %d\n", ProductOfDepartureField(rules, myticket, nearbies))
}
