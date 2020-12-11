package main

func getNeighbors(rows []string, x, y int) []string {
	deltas := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	var ngh []string
	for _, d := range deltas {
		xd := x + d[0]
		yd := y + d[1]
		if xd >= 0 && xd < len(rows[0]) && yd >= 0 && yd < len(rows) {
			ngh = append(ngh, getSeat(rows, xd, yd))
		}
	}
	return ngh
}

func countOccupiedNeighbors(rows []string, x, y int) (fs int) {
	for _, v := range getNeighbors(rows, x, y) {
		if v == occupied {
			fs++
		}
	}
	return fs
}

func getNewIterationTask1(rows []string) []string {
	var it []string
	for y, col := range rows {
		line := ""
		for x := range col {
			switch {
			case getSeat(rows, x, y) == empty && countOccupiedNeighbors(rows, x, y) == 0:
				line = line + occupied
				break
			case getSeat(rows, x, y) == occupied && countOccupiedNeighbors(rows, x, y) >= 4:
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

func task1(rows []string) int {
	last := rows

	i := getNewIterationTask1(last)
	for {
		last = i
		i = getNewIterationTask1(last)

		//	fmt.Printf(" %+v\n %+v %t\n\n", last, i, compare(last, i))
		if compare(last, i) {
			break
		}
	}
	//fmt.Printf(" %+v\n", last)

	return countOccupiedSeats(last)
}
