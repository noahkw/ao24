package main

import (
	"common"
	"fmt"
	"time"
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
const VISITED_UP = 'X'
const VISITED_DOWN = 'Y'
const VISITED_RIGHT = 'Z'
const VISITED_LEFT = 'L'
const OUTSIDE = '?'
const GUARD = '^'

func isVisited(tile rune) bool {
	return tile == VISITED_RIGHT || tile == VISITED_LEFT || tile == VISITED_DOWN || tile == VISITED_UP
}

func dirToVisited(dir Direction) rune {
	switch dir {
	case Up:
		return VISITED_UP
	case Down:
		return VISITED_DOWN
	case Right:
		return VISITED_RIGHT
	case Left:
		return VISITED_LEFT
	default:
		panic("unknown direction" + dir)
	}
}

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

// returns whether the sim loops
func simulate(lines []string, obstaclePos Position) bool {
	currentTile := getTileAt(&lines, obstaclePos)
	if currentTile != WALKABLE {
		return false
	}

	linesCopy := make([]string, len(lines))
	copy(linesCopy, lines)

	setTileAt(&linesCopy, obstaclePos, OBSTACLE)
	currentPosition, currentDir := getGuardPosition(&linesCopy)
	setTileAt(&linesCopy, currentPosition, dirToVisited(currentDir))

	for {
		newPos := moveInDir(currentPosition, currentDir)
		newTile := getTileAt(&linesCopy, newPos)

		if newTile == WALKABLE || isVisited(newTile) {
			currentPosition = newPos
			setTileAt(&linesCopy, currentPosition, dirToVisited(currentDir))
		} else if newTile == OBSTACLE {
			currentDir = nextDir(currentDir)
		} else if newTile == OUTSIDE {
			return false
		}

		if newTile == dirToVisited(currentDir) {
			return true
		}
	}
}

func main() {
	start := time.Now()
	lines := common.ReadLinesFromFile("src/06/input.txt")

	fmt.Println(string(getTileAt(&lines, Position{x: -1, y: 0})))

	numLoopingSims := 0

	for i := 0; i < len(lines)-1; i++ {
		for k := 0; k < len((lines)[0])-1; k++ {
			obstaclePos := Position{x: i, y: k}

			isLooping := simulate(lines, obstaclePos)

			if isLooping {
				numLoopingSims++
			}
		}
	}

	fmt.Printf("looping: %d\n", numLoopingSims)
	fmt.Printf("execution time: %v\n", time.Since(start))
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
