package main

import (
	"common"
	"fmt"
)

func isSafeStep(a int, b int) bool {
	diff := common.Abs(a - b)

	return diff >= 1 && diff <= 3
}

func removeIndexFromSlice(slice []int, idx int) []int {
	out := make([]int, len(slice))
	copy(out, slice)

	newSlice := append(out[:idx], out[idx+1:]...)

	return newSlice
}

func checkReportVariations(report []int, isIncreasing bool) bool {
	for i := -1; i < len(report); i++ {
		if checkSingleReport(report, isIncreasing, i) {
			return true
		}
	}

	return false
}

func checkSingleReport(report []int, isIncreasing bool, leaveOutIndex int) bool {
	var newSlice []int

	if leaveOutIndex > -1 {
		newSlice = removeIndexFromSlice(report, leaveOutIndex)
	} else {
		newSlice = report
	}

	for i := 0; i < len(newSlice)-1; i++ {
		tokenA := newSlice[i]
		tokenB := newSlice[i+1]

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

		isSafe := checkReportVariations(tokens, true) || checkReportVariations(tokens, false)

		if isSafe {
			numSafe += 1
		}

		fmt.Printf("%v [safe: %t]\n", tokens, isSafe)
	}

	fmt.Printf("%d reports are safe", numSafe)
}
