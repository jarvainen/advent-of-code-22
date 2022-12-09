package __day_nine

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content embed.FS

type void struct{}

var member void

func intAbs(x int) float64 {
	return math.Abs(float64(x))
}

func sign(n int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	return 1
}

func silver() {
	visitedPositions := make(map[string]void)
	headX, headY := 0, 0
	tailX, tailY := 0, 0
	util.ReadFileByLine(content, "input.txt", func(line string) {
		data := strings.Fields(line)
		direction := data[0]
		steps, _ := strconv.Atoi(data[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				headX += 1
				break
			case "L":
				headX -= 1
				break
			case "U":
				headY -= 1
				break
			case "D":
				headY += 1
				break
			}

			xDiff := headX - tailX
			yDiff := headY - tailY
			if intAbs(xDiff) > 1 || intAbs(yDiff) > 1 {
				tailX += sign(xDiff)
				tailY += sign(yDiff)
			}
			visitedPositions[fmt.Sprintf("%d:%d", tailX, tailY)] = member
		}
	})
	solution := len(visitedPositions)
	fmt.Printf("Solution is %d\n", solution)
}

type xy struct {
	x int
	y int
}

func gold() {
	visitedPositions := make(map[string]void)
	headX, headY := 0, 0
	knots := make([]xy, 10)
	util.ReadFileByLine(content, "input.txt", func(line string) {
		data := strings.Fields(line)
		direction := data[0]
		steps, _ := strconv.Atoi(data[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "R":
				headX += 1
				break
			case "L":
				headX -= 1
				break
			case "U":
				headY -= 1
				break
			case "D":
				headY += 1
				break
			}

			for j := 0; j < len(knots); j++ {
				prevX := headX
				prevY := headY
				if j > 0 {
					prevX = knots[j-1].x
					prevY = knots[j-1].y
				}
				xDiff := prevX - knots[j].x
				yDiff := prevY - knots[j].y
				if intAbs(xDiff) > 1 || intAbs(yDiff) > 1 {
					knots[j].x += sign(xDiff)
					knots[j].y += sign(yDiff)
				}
			}

			visitedPositions[fmt.Sprintf("%d:%d", knots[8].x, knots[8].y)] = member
		}
	})
	solution := len(visitedPositions)
	fmt.Printf("Solution is %d\n", solution)
}

func Solution() {
	silver()
	gold()
}
