package main

import (
	"common"
	"errors"
	"fmt"
)

func getCharAt(pos Position, lines []string) (uint8, error) {
	if pos.x > len(lines[0])-1 || pos.x < 0 || pos.y < 0 || pos.y > len(lines)-1 {
		return 0, errors.New("no")
	}

	return lines[pos.y][pos.x], nil
}

func searchInDirection(toSearch string, pos Position, deltaX, deltaY int, lines []string) bool {
	for i := 0; i < len(toSearch); i++ {
		searchChar := toSearch[i]
		charToCompare, err := getCharAt(Position{x: pos.x + i*deltaX, y: pos.y + i*deltaY}, lines)
		if searchChar != charToCompare || err != nil {
			return false
		}
	}
	return true
}

type Position struct {
	x int
	y int
}

func searchAtPosition(toSearch string, lines []string) int {
	directions := []struct {
		dx int
		dy int
	}{
		{1, 0},   // right
		{-1, 0},  // left
		{0, -1},  // up
		{0, 1},   // down
		{1, -1},  // up-right
		{-1, -1}, // up-left
		{1, 1},   // down-right
		{-1, 1},  // down-left
	}

	numFound := 0
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			pos := Position{x: x, y: y}
			for _, dir := range directions {
				if searchInDirection(toSearch, pos, dir.dx, dir.dy, lines) {
					numFound++
				}
			}
		}
	}
	return numFound
}

func main() {
	searchString := "XMAS"

	//lines := common.ReadLinesFromFile("src/04/testinput.txt")
	lines := common.ReadLinesFromFile("src/04/input.txt")
	fmt.Println(lines)

	fmt.Println(searchAtPosition(searchString, lines))
}
