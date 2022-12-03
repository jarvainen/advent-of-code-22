package __day_three

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"strings"
)

//go:embed input.txt
var content embed.FS

func getItemTypePriority(itemType string) int {
	typeListUpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	typeListLowerCase := strings.ToLower(typeListUpperCase)
	typeList := typeListLowerCase + typeListUpperCase
	return strings.Index(typeList, itemType) + 1
}

func findStringIntersection(first string, second string) (result string) {
	for _, firstCharacter := range first {
		for _, secondCharacter := range second {
			if firstCharacter == secondCharacter {
				return fmt.Sprintf("%c", firstCharacter)
			}
		}
	}
	return ""
}

func findThreeStringIntersection(first string, second string, third string) (result string) {
	for _, firstCharacter := range first {
		for _, secondCharacter := range second {
			for _, thirdCharacter := range third {
				if firstCharacter == secondCharacter && firstCharacter == thirdCharacter {
					return fmt.Sprintf("%c", firstCharacter)
				}
			}
		}
	}
	return ""
}

func silver() {
	totalPriority := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		halfIndex := len(line) / 2
		firstCompartment := line[:halfIndex]
		secondCompartment := line[halfIndex:]
		compartmentType := findStringIntersection(firstCompartment, secondCompartment)
		totalPriority += getItemTypePriority(compartmentType)
	})
	fmt.Printf("Solution is %d\n", totalPriority)
}

func gold() {
	totalPriority := 0
	var badgeGroup []string
	util.ReadFileByLine(content, "input.txt", func(line string) {
		if len(badgeGroup) < 3 {
			badgeGroup = append(badgeGroup, line)
		}
		if len(badgeGroup) == 3 {
			badgeType := findThreeStringIntersection(badgeGroup[0], badgeGroup[1], badgeGroup[2])
			totalPriority += getItemTypePriority(badgeType)
			badgeGroup = make([]string, 0)
		}

	})
	fmt.Printf("Solution is %d\n", totalPriority)
}

func Solution() {
	silver()
	gold()
}
