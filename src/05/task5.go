package main

import (
	"common"
	"fmt"
	"strconv"
	"strings"
)

// PageOrder An ordering rule for two pages
// ex: 47|53 -> before=47, after=53
type PageOrder struct {
	// the page that has to come first
	before int
	// the page that has to come after but not necessarily directly
	after int
}

type Update struct {
	pageNums []int
	pageSet  map[int]bool
}

func parseUpdate(in string) Update {
	tokens := strings.Split(in, ",")

	if len(tokens) == 0 {
		panic("input could not be parsed: " + in)
	}

	out := make([]int, 0)
	outSet := make(map[int]bool)

	for _, token := range tokens {
		pageNum, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}
		out = append(out, pageNum)
		outSet[pageNum] = true
	}

	return Update{pageNums: out, pageSet: outSet}
}

func (update Update) getMiddlePage() int {
	if len(update.pageNums)%2 != 1 {
		panic("cannot get middle page")
	}

	return update.pageNums[len(update.pageNums)/2]
}

func (update Update) checkUpdateAtIndex(idx int, pageOrders *[]PageOrder) (bool, *int, *PageOrder) {
	pageToCheck := update.pageNums[idx]

	for _, pageOrder := range *pageOrders {
		if pageOrder.after != pageToCheck {
			// irrelevant
			// maybe check pageOrder.before too
			continue
		}

		for i := idx + 1; i < len(update.pageNums); i++ {
			curPage := update.pageNums[i]

			if curPage == pageOrder.before {
				fmt.Printf("violated PageOrder [%d, %d]", pageOrder.before, pageOrder.after)
				return false, &i, &pageOrder
			}
		}
	}

	return true, nil, nil
}

func (update Update) fixBadUpdate(pageOrders *[]PageOrder) Update {
	fixedUpdate := update

	swapIndices := make([]int, 0)
	for idx := 0; idx < len(fixedUpdate.pageNums); idx++ {
		goodIndex, indexToSwap, pageOrder := fixedUpdate.checkUpdateAtIndex(idx, pageOrders)
		_ = pageOrder
		if !goodIndex {
			swapIndices = append(swapIndices, idx)
			page := fixedUpdate.pageNums[idx]
			pageToSwap := fixedUpdate.pageNums[*indexToSwap]
			fixedUpdate.pageNums[idx] = pageToSwap
			fixedUpdate.pageNums[*indexToSwap] = page
		}
	}

	return fixedUpdate
}

func (update Update) fixUntilGood(pageOrders *[]PageOrder) Update {
	for {
		update.fixBadUpdate(pageOrders)

		allIndicesGood := true
		for idx := 0; idx < len(update.pageNums); idx++ {
			goodIndex, _, _ := update.checkUpdateAtIndex(idx, pageOrders)

			if !goodIndex {
				allIndicesGood = false
				break
			}
		}

		if allIndicesGood {
			return update
		}
	}
}

func parsePageOrder(in string) PageOrder {
	tokens := strings.Split(in, "|")

	if len(tokens) != 2 {
		panic("input could not be parsed: " + in)
	}

	before, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	after, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic(err)
	}

	return PageOrder{
		before: before,
		after:  after,
	}
}

func main() {
	//lines := common.ReadLinesFromFile("src/05/testinput.txt")
	lines := common.ReadLinesFromFile("src/05/input.txt")
	fmt.Println(lines)

	parsingStage := 0

	pageOrders := make([]PageOrder, 0)
	updates := make([]Update, 0)

	for _, line := range lines {
		if len(line) == 0 {
			fmt.Println("done parsing PageOrders")
			parsingStage++
			continue
		}

		if parsingStage == 0 {
			pageOrders = append(pageOrders, parsePageOrder(line))
		} else if parsingStage == 1 {
			updates = append(updates, parseUpdate(line))
		} else {
			panic("how did we get here?")
		}
	}

	fmt.Println("PageOrders")
	fmt.Println(pageOrders)
	fmt.Println("updates")
	fmt.Println(updates)

	middlePageSum := 0
	badUpdates := make([]Update, 0)

	for _, update := range updates {
		for i := 0; i < len(update.pageNums); i++ {
			isGood, _, _ := update.checkUpdateAtIndex(i, &pageOrders)
			if !isGood {
				badUpdates = append(badUpdates, update)
				break
			}
		}
	}

	fmt.Println("bad updates\n\n")
	fmt.Println(badUpdates)

	for _, badUpdate := range badUpdates {
		fixedUpdate := badUpdate.fixUntilGood(&pageOrders)
		fmt.Printf("\nbad update: %s\nfixed update: %s\n", badUpdate, fixedUpdate)

		middlePageSum += fixedUpdate.getMiddlePage()
	}

	fmt.Printf("\n---\nmiddle page sum: %d", middlePageSum)
}
