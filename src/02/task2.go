package main

import (
	"common"
	"fmt"
)

func isSafeStep(a int, b int) bool {
	diff := common.Abs(a - b)

	return diff >= 1 && diff <= 3
}

func checkReportInc(report []int, isIncreasing bool) bool {
	for i := 0; i < len(report)-1; i++ {
		tokenA := report[i]
		tokenB := report[i+1]

		var isNotMonotonous bool

		if isIncreasing {
			isNotMonotonous = tokenB >= tokenA
		} else {
			isNotMonotonous = tokenA >= tokenB
		}

		if isNotMonotonous || !isSafeStep(tokenA, tokenB) {
			return false
		}
	}

	return true
}

func main() {
	lines := common.ReadLinesFromFile("src/02/input.txt")

	numSafe := 0

	for _, line := range lines {
		tokens := common.TokenizeLineAsInts(line)

		isSafe := checkReportInc(tokens, true) || checkReportInc(tokens, false)

		if isSafe {
			numSafe += 1
		}

		fmt.Printf("%v [safe: %t]\n", tokens, isSafe)
	}

	fmt.Printf("%d reports are safe", numSafe)
}
