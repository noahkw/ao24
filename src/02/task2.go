package main

import (
	"common"
	"fmt"
)

func isSafeStep(a int, b int) bool {
	diff := common.Abs(a - b)

	return diff >= 1 && diff <= 3
}

func main() {
	fmt.Println(isSafeStep(2, 4))
}
