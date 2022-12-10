package _0_day_ten

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
	x := 1
	cycles := 0
	solution := 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		fields := strings.Fields(line)
		cmd := fields[0]

		cycleDuration := 1
		if cmd == "addx" {
			cycleDuration = 2
		}
		for i := 0; i < cycleDuration; i++ {
			cycles += 1
			if cycles%40 == 20 {
				solution += cycles * x
			}
		}

		if cmd == "addx" {
			value, _ := strconv.Atoi(fields[1])
			x += value
		}
	})
	fmt.Printf("Solution is %d\n", solution)
}

func gold() {
	x := 1
	cycles := 0
	row := 0
	fmt.Println("Solution is")
	util.ReadFileByLine(content, "input.txt", func(line string) {
		fields := strings.Fields(line)
		cmd := fields[0]

		cycleDuration := 1
		if cmd == "addx" {
			cycleDuration = 2
		}
		for i := 0; i < cycleDuration; i++ {
			cycles += 1
			rowCycle := cycles - (row * 40) - 1
			if rowCycle >= x-1 && rowCycle <= x+1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if cycles%40 == 0 && cycles != 0 {
				fmt.Println()
				row += 1
			}
		}

		if cmd == "addx" {
			value, _ := strconv.Atoi(fields[1])
			x += value
		}
	})

}

func Solution() {
	silver()
	gold()
}
