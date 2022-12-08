package __day_eight

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
)

//go:embed input.txt
var content embed.FS

func isVisible(grid [][]int, x int, y int) bool {
	if x == 0 || x == len(grid[0])-1 || y == 0 || y == len(grid)-1 {
		return true
	}

	visibleFromTop := true
	tree := grid[y][x]
	for i := 0; i < y; i++ {
		currentTree := grid[i][x]
		if tree <= currentTree {
			visibleFromTop = false
		}
	}

	visibleFromBottom := true
	for i := y + 1; i < len(grid); i++ {
		currentTree := grid[i][x]
		if tree <= currentTree {
			visibleFromBottom = false
		}
	}

	visibleFromRight := true
	for i := x + 1; i < len(grid[0]); i++ {
		currentTree := grid[y][i]
		if tree <= currentTree {
			visibleFromRight = false
		}
	}

	visibleFromLeft := true
	for i := 0; i < x; i++ {
		currentTree := grid[y][i]
		if tree <= currentTree {
			visibleFromLeft = false
		}
	}

	return visibleFromRight || visibleFromLeft || visibleFromTop || visibleFromBottom
}

func silver() {
	grid := make([][]int, 0)
	util.ReadFileByLine(content, "input.txt", func(line string) {
		row := make([]int, len(line))
		for x, r := range line {
			row[x] = int(r - '0')
		}
		grid = append(grid, row)
	})

	solution := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if isVisible(grid, x, y) {
				solution += 1
			}
		}
	}

	fmt.Printf("Solution is %d\n", solution)
}

func getSceningScore(grid [][]int, x int, y int) int {
	tree := grid[y][x]

	viewDistanceTop := 0
	for i := y - 1; i >= 0; i-- {
		currentTree := grid[i][x]
		viewDistanceTop += 1
		if currentTree >= tree {
			break
		}
	}

	viewDistanceBottom := 0
	for i := y + 1; i < len(grid); i++ {
		currentTree := grid[i][x]
		viewDistanceBottom += 1
		if currentTree >= tree {
			break
		}
	}

	viewDistanceRight := 0
	for i := x + 1; i < len(grid[0]); i++ {
		currentTree := grid[y][i]
		viewDistanceRight += 1
		if currentTree >= tree {
			break
		}
	}

	viewDistanceLeft := 0
	for i := x - 1; i >= 0; i-- {
		currentTree := grid[y][i]
		viewDistanceLeft += 1
		if currentTree >= tree {
			break
		}
	}

	return viewDistanceRight * viewDistanceBottom * viewDistanceTop * viewDistanceLeft
}

func gold() {
	grid := make([][]int, 0)
	util.ReadFileByLine(content, "input.txt", func(line string) {
		row := make([]int, len(line))
		for x, r := range line {
			row[x] = int(r - '0')
		}
		grid = append(grid, row)
	})

	solution := -1
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			scenicScore := getSceningScore(grid, x, y)
			if scenicScore > solution {
				solution = scenicScore
			}
		}
	}

	fmt.Printf("Solution is %d\n", solution)
}

func Solution() {
	silver()
	gold()
}
