package __day_four

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"strconv"
	"strings"
)

//go:embed input.txt
var content embed.FS

func silver() {
	fullContainedCount := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		pairs := strings.Split(line, ",")
		firstRange := strings.Split(pairs[0], "-")
		secondRange := strings.Split(pairs[1], "-")
		firstSectionStart, _ := strconv.Atoi(firstRange[0])
		firstSectionEnd, _ := strconv.Atoi(firstRange[1])
		secondSectionStart, _ := strconv.Atoi(secondRange[0])
		secondSectionEnd, _ := strconv.Atoi(secondRange[1])
		firstContainsSecond := secondSectionStart >= firstSectionStart && secondSectionEnd <= firstSectionEnd
		secondContainsFirst := firstSectionStart >= secondSectionStart && firstSectionEnd <= secondSectionEnd
		if firstContainsSecond || secondContainsFirst {
			fullContainedCount += 1
		}
	})
	fmt.Printf("Solution is %d\n", fullContainedCount)
}

func gold() {
	overlap := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		pairs := strings.Split(line, ",")
		firstRange := strings.Split(pairs[0], "-")
		secondRange := strings.Split(pairs[1], "-")
		firstSectionStart, _ := strconv.Atoi(firstRange[0])
		firstSectionEnd, _ := strconv.Atoi(firstRange[1])
		secondSectionStart, _ := strconv.Atoi(secondRange[0])
		secondSectionEnd, _ := strconv.Atoi(secondRange[1])
		isOverlapping := firstSectionEnd >= secondSectionStart && firstSectionStart <= secondSectionEnd
		if isOverlapping {
			overlap += 1
		}
	})
	fmt.Printf("Solution is %d\n", overlap)
}

func Solution() {
	silver()
	gold()
}
