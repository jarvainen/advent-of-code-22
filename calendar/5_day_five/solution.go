package __day_five

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"regexp"
	"strconv"
)

// move 1 from 3 to 9

//go:embed input.txt
var content embed.FS

func silver() {
	stacks := make([][]string, 9)
	initRowCount := 0
	re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
	util.ReadFileByLine(content, "input.txt", func(line string) {
		if line == "" || line[0:2] == " 1" {
			return
		}
		if initRowCount <= 7 {
			stackIndex := 0
			for i := 1; i < len(line)-1; i += 4 {
				crate := line[i : i+1]
				if crate != " " {
					stacks[stackIndex] = append([]string{crate}, stacks[stackIndex]...)
				}
				stackIndex += 1
			}
			initRowCount += 1
			return
		}

		match := re.FindStringSubmatch(line)
		count, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])

		for i := 0; i < count; i++ {
			fromIndex := from - 1
			toIndex := to - 1
			fromStack := stacks[fromIndex]
			var crate string
			var newStack []string
			if len(fromStack) > 1 {
				crate, newStack = fromStack[len(fromStack)-1], fromStack[:len(fromStack)-1]
			} else {
				crate, newStack = fromStack[len(fromStack)-1], make([]string, 0)
			}
			stacks[toIndex] = append(stacks[toIndex], crate)
			stacks[fromIndex] = newStack
		}
	})
	solution := ""
	for _, stack := range stacks {
		solution += stack[len(stack)-1]
	}
	fmt.Printf("Solution is %s\n", solution)
}

func gold() {
	stacks := make([][]string, 9)
	initRowCount := 0
	re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
	util.ReadFileByLine(content, "input.txt", func(line string) {
		if line == "" || line[0:2] == " 1" {
			return
		}
		if initRowCount <= 7 {
			stackIndex := 0
			for i := 1; i < len(line)-1; i += 4 {
				crate := line[i : i+1]
				if crate != " " {
					stacks[stackIndex] = append([]string{crate}, stacks[stackIndex]...)
				}
				stackIndex += 1
			}
			initRowCount += 1
			return
		}

		match := re.FindStringSubmatch(line)
		count, _ := strconv.Atoi(match[1])
		from, _ := strconv.Atoi(match[2])
		to, _ := strconv.Atoi(match[3])

		fromIndex := from - 1
		toIndex := to - 1
		fromStack := stacks[fromIndex]
		var crates []string
		var newStack []string
		if len(fromStack) > 1 {
			crates, newStack = fromStack[len(fromStack)-count:], fromStack[:len(fromStack)-count]
		} else {
			crates, newStack = fromStack[len(fromStack)-count:], make([]string, 0)
		}
		stacks[toIndex] = append(stacks[toIndex], crates...)
		stacks[fromIndex] = newStack
	})
	solution := ""
	for _, stack := range stacks {
		solution += stack[len(stack)-1]
	}
	fmt.Printf("Solution is %s\n", solution)
}

func Solution() {
	silver()
	gold()
}
