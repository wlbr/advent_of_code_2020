package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var input = "input.txt"

func readDataset(fname string) (data []int) {

	f, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", input, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		num, e := strconv.Atoi(line)
		if e != nil {
			log.Fatalf("Not a number in dataset '%s':  %s", line, e)
		}
		data = append(data, num)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return data
}

func search2(sum int, candidate int, data []int) (tupel [][]int) {
	for _, d := range data {
		if candidate+d == sum {
			tupel = append(tupel, []int{candidate, d})
		}
	}
	return tupel
}

func search3(sum int, candidate int, data []int) (tupel [][]int) {
	for i, c := range data {
		for _, d := range data[i+1:] {
			if candidate+c+d == sum {
				tupel = append(tupel, []int{candidate, c, d})
			}
		}
	}
	return tupel
}

func main() {
	d := readDataset(input)

	var res2 [][]int
	for i, n := range d {
		res2 = append(res2, search2(2020, n, d[i+1:])...)
	}
	for _, t := range res2 {
		fmt.Printf("2 components: %d\n", t[0]*t[1])
	}

	var res3 [][]int
	for i, n := range d {
		res3 = append(res3, search3(2020, n, d[i+1:])...)
	}
	for _, t := range res3 {
		fmt.Printf("3 components: %d\n", t[0]*t[1]*t[2])
	}
}
