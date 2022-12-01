package __day_one

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"log"
	"strconv"
)

//go:embed input.txt
var content embed.FS

func silver() {
	largestCapacity := -1
	currentElfTotal := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		if line == "" {
			if currentElfTotal > largestCapacity {
				largestCapacity = currentElfTotal
			}
			currentElfTotal = 0
			return
		}

		lineAsInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		currentElfTotal += lineAsInt
	})
	fmt.Printf("Solution is %d\n", largestCapacity)
}

func sortedInsert(array []int, value int) []int {
	for i, num := range array {
		if num < value {
			replaced := array[i]
			array[i] = value
			return sortedInsert(array, replaced)
		}
	}
	return array
}

func gold() {
	elvesCarrying := []int{-1, -1, -1}
	currentElfTotal := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		if line == "" {
			elvesCarrying = sortedInsert(elvesCarrying, currentElfTotal)
			currentElfTotal = 0
			return
		}

		lineAsInt, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		currentElfTotal += lineAsInt
	})
	allTotal := 0
	for _, value := range elvesCarrying {
		allTotal += value
	}
	fmt.Printf("Solution is %d\n", allTotal)
}

func Solution() {
	silver()
	gold()
}
