package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func TokenizeLine(line string) []string {
	return strings.Fields(line)
}

func TokenizeLineAsInts(line string) []int {
	tokens := TokenizeLine(line)

	var out []int

	for _, token := range tokens {
		intToken, err := strconv.Atoi(token)

		if err != nil {
			panic(err)
		}

		out = append(out, intToken)
	}

	return out
}

func ReadLinesFromFile(filePath string) []string {
	// open the file
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("error opening file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error scanning file", err)
	}

	return lines
}
