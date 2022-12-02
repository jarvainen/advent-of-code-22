package __day_two

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"strings"
)

//go:embed input.txt
var content embed.FS

var meRockPaperScissorValues = map[string]int{
	"X": 1, // Rock
	"Y": 2, // Paper
	"Z": 3, // Scissors
}

var opponentRockPaperScissorValues = map[string]int{
	"A": 1, // Rock
	"B": 2, // Paper
	"C": 3, // Scissors
}

func processMatch(me int, opponent int) int {
	roundScore := 0

	winning := (me == 1 && opponent == 3) || (me == 2 && opponent == 1) || (me == 3 && opponent == 2)

	roundScore += me
	if me == opponent {
		roundScore += 3
	} else if winning {
		roundScore += 6
	}
	return roundScore
}

func silver() {
	totalScore := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		splittedString := strings.Fields(line)
		opponentString := splittedString[0]
		opponentValue := opponentRockPaperScissorValues[opponentString]
		meString := splittedString[1]
		meValue := meRockPaperScissorValues[meString]

		totalScore += processMatch(meValue, opponentValue)
	})
	fmt.Printf("Solution is %d\n", totalScore)
}

func getLoserValue(opponentValue int) int {
	if opponentValue == 1 {
		return 3
	} else if opponentValue == 2 {
		return 1
	} else if opponentValue == 3 {
		return 2
	}
	return 0
}

func getWinnerValue(opponentValue int) int {
	if opponentValue == 1 {
		return 2
	} else if opponentValue == 2 {
		return 3
	} else if opponentValue == 3 {
		return 1
	}
	return 0
}

func gold() {
	totalScore := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		splittedString := strings.Fields(line)
		opponentString := splittedString[0]
		opponentValue := opponentRockPaperScissorValues[opponentString]
		meString := splittedString[1]

		meValue := 0
		switch meString {
		case "X":
			meValue = getLoserValue(opponentValue)
			break
		case "Y":
			meValue = opponentValue
			break
		case "Z":
			meValue = getWinnerValue(opponentValue)
			break
		}

		totalScore += processMatch(meValue, opponentValue)
	})
	fmt.Printf("Solution is %d\n", totalScore)
}

func Solution() {
	silver()
	gold()
}
