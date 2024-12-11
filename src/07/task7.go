package main

import (
	"common"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Equation struct {
	result   int
	operands []int
}

func (equation Equation) compute() int {
	num := len(equation.operands) - 1
	maxCount := int(math.Pow(2, float64(num)))

	for i := range maxCount {
		combination := strconv.FormatInt(int64(i), 2)
		combinationString := fmt.Sprintf("%0*s", num, combination)

		result := equation.tryCombination(combinationString)

		if result == equation.result {
			return result
		}
	}

	return 0
}

func (equation Equation) tryCombination(combinationString string) int {
	result := equation.operands[0]
	for index, char := range combinationString {
		nextOperand := equation.operands[index+1]

		if char == '0' {
			result += nextOperand
		} else if char == '1' {
			result *= nextOperand
		}
	}
	return result
}

func parseInt(in string) int {
	result, err := strconv.Atoi(in)

	if err != nil {
		panic(err)
	}

	return result
}

func parseEquation(in string) Equation {
	tokens := strings.Fields(in)
	tokenInts := make([]int, 0)

	for _, token := range tokens[1:] {
		tokenInts = append(tokenInts, parseInt(token))
	}

	return Equation{result: parseInt(tokens[0][:len(tokens[0])-1]), operands: tokenInts}
}

func main() {
	lines := common.ReadLinesFromFile("src/07/input.txt")

	possibleEquationSum := 0
	for _, line := range lines {
		equation := parseEquation(line)

		possibleEquationSum += equation.compute()
	}

	fmt.Printf("sum: %d", possibleEquationSum)
}
