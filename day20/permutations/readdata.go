package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadData(fname string) (tiles []*Tile) {

	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var currentTile string
	var currentFields []string
	for scanner.Scan() {
		line := strings.Trim(strings.ToLower(scanner.Text()), " ")
		if strings.HasPrefix(line, "tile") {
			currentTile = line
		} else if line == "" {
			tiles = append(tiles, NewTile(currentTile, currentFields))
			currentTile = ""
			currentFields = []string{}
		} else {
			currentFields = append(currentFields, line)
		}
	}
	tiles = append(tiles, NewTile(currentTile, currentFields))

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return tiles
}
