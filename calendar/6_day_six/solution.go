package __day_six

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
)

//go:embed input.txt
var content embed.FS

func hasAllUniqueCharacters(str string) bool {
	charMap := map[rune]bool{}
	for _, char := range str {
		_, ok := charMap[char]
		if ok {
			return false
		}
		charMap[char] = true
	}
	return true
}

func findMarker(stream string, consecutiveLength int) int {
	recentCharacters := ""
	for index, char := range stream {
		if len(recentCharacters) == consecutiveLength {
			if hasAllUniqueCharacters(recentCharacters) {
				return index
			} else {
				recentCharacters = recentCharacters[1:]
			}
		}
		recentCharacters += fmt.Sprintf("%c", char)
	}
	return -1
}

func silver() {
	solution := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		solution = findMarker(line, 4)
	})
	fmt.Printf("Solution is %d\n", solution)
}

func gold() {
	solution := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		solution = findMarker(line, 14)
	})
	fmt.Printf("Solution is %d\n", solution)
}

func Solution() {
	silver()
	gold()
}
