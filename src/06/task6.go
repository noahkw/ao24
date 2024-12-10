package main

import (
	"common"
	"fmt"
)

type Direction string

const (
	Up    Direction = "Up"
	Down  Direction = "Down"
	Right Direction = "Right"
	Left  Direction = "Left"
)

const OBSTACLE = '#'
const WALKABLE = '.'
const VISITED = 'X'
const OUTSIDE = '?'
const GUARD = '^'

func nextDir(dir Direction) Direction {
	switch dir {
	case Up:
		return Right
	case Right:
		return Down
	case Down:
		return Left
	case Left:
		return Up
	default:
		panic("unknown direction" + dir)
	}
}

type Position struct {
	x int
	y int
}

func getTileAt(lines *[]string, pos Position) rune {
	if pos.y < 0 || pos.y > len(*lines)-1 || pos.x < 0 || pos.x > len((*lines)[0])-1 {
		return OUTSIDE
	}

	character := (*lines)[pos.y][pos.x]
	return rune(character)
}

func setTileAt(lines *[]string, pos Position, newTile rune) {
	(*lines)[pos.y] = (*lines)[pos.y][:pos.x] + string(newTile) + (*lines)[pos.y][pos.x+1:]
}

func moveInDir(pos Position, dir Direction) Position {
	switch dir {
	case Up:
		return Position{x: pos.x, y: pos.y - 1}
	case Down:
		return Position{x: pos.x, y: pos.y + 1}
	case Right:
		return Position{x: pos.x + 1, y: pos.y}
	case Left:
		return Position{x: pos.x - 1, y: pos.y}
	default:
		panic("unknown direction" + dir)
	}
}

func main() {
	lines := common.ReadLinesFromFile("src/06/input.txt")

	fmt.Println(string(getTileAt(&lines, Position{x: -1, y: 0})))

	var currentTile rune

	currentPosition, currentDir := getGuardPosition(&lines)
	fmt.Printf("found guard at %d with direction %s", currentPosition, currentDir)
	setTileAt(&lines, currentPosition, VISITED)

	for currentTile != OUTSIDE {
		newPos := moveInDir(currentPosition, currentDir)
		newTile := getTileAt(&lines, newPos)

		if newTile == WALKABLE || newTile == VISITED {
			currentPosition = newPos
			setTileAt(&lines, currentPosition, VISITED)
		} else if newTile == OBSTACLE {
			currentDir = nextDir(currentDir)
		} else if newTile == OUTSIDE {
			break
		}
	}

	fmt.Printf("visited tiles: %d", countVisited(&lines))
}

func getGuardDir(guardChar rune) Direction {
	if guardChar == '^' {
		return Up
	} else if guardChar == 'v' {
		return Down
	} else if guardChar == '<' {
		return Left
	} else if guardChar == '>' {
		return Right
	} else {
		return ""
	}
}

func getGuardPosition(lines *[]string) (Position, Direction) {
	for i := 0; i < len(*lines); i++ {
		for k := 0; k < len((*lines)[0]); k++ {
			curPos := Position{x: i, y: k}

			guardDir := getGuardDir(getTileAt(lines, curPos))
			if guardDir != "" {
				return curPos, guardDir
			}
		}
	}

	panic("could not find guard")
}

func countVisited(lines *[]string) int {
	count := 0
	for i := 0; i < len(*lines); i++ {
		for k := 0; k < len((*lines)[0]); k++ {
			if getTileAt(lines, Position{x: i, y: k}) == VISITED {
				count++
			}
		}
	}
	return count
}
