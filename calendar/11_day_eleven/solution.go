package _1_day_eleven

import (
	"embed"
	"fmt"
)

//go:embed input.txt
var content embed.FS

type monkey struct {
	items     []int
	operation func(int) int
	test      func(int) bool
	testTrue  int
	testFalse int
}

func createMonkey(data monkey) *monkey {
	return &data
}

var monkeys = []*monkey{}

func initMonkeys() {
	monkeys = []*monkey{
		createMonkey(monkey{
			items: []int{
				63,
				84,
				80,
				83,
				84,
				53,
				88,
				72,
			},
			operation: func(item int) int {
				return item * 11
			},
			test: func(item int) bool {
				return item%13 == 0
			},
			testTrue:  4,
			testFalse: 7,
		}),
		createMonkey(monkey{
			items: []int{
				67,
				56,
				92,
				88,
				84,
			},
			operation: func(item int) int {
				return item + 4
			},
			test: func(item int) bool {
				return item%11 == 0
			},
			testTrue:  5,
			testFalse: 3,
		}),
		createMonkey(monkey{
			items: []int{
				52,
			},
			operation: func(item int) int {
				return item * item
			},
			test: func(item int) bool {
				return item%2 == 0
			},
			testTrue:  3,
			testFalse: 1,
		}),
		createMonkey(monkey{
			items: []int{
				59,
				53,
				60,
				92,
				69,
				72,
			},
			operation: func(item int) int {
				return item + 2
			},
			test: func(item int) bool {
				return item%5 == 0
			},
			testTrue:  5,
			testFalse: 6,
		}),
		createMonkey(monkey{
			items: []int{
				61,
				52,
				55,
				61,
			},
			operation: func(item int) int {
				return item + 3
			},
			test: func(item int) bool {
				return item%7 == 0
			},
			testTrue:  7,
			testFalse: 2,
		}),
		createMonkey(monkey{
			items: []int{
				79,
				53,
			},
			operation: func(item int) int {
				return item + 1
			},
			test: func(item int) bool {
				return item%3 == 0
			},
			testTrue:  0,
			testFalse: 6,
		}),
		createMonkey(monkey{
			items: []int{
				59,
				86,
				67,
				95,
				92,
				77,
				91,
			},
			operation: func(item int) int {
				return item + 5
			},
			test: func(item int) bool {
				return item%19 == 0
			},
			testTrue:  4,
			testFalse: 0,
		}),
		createMonkey(monkey{
			items: []int{
				58,
				83,
				89,
			},
			operation: func(item int) int {
				return item * 19
			},
			test: func(item int) bool {
				return item%17 == 0
			},
			testTrue:  2,
			testFalse: 1,
		}),
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func both() {
	roundData := [2]int{
		20,
		10000,
	}

	lcm := LCM(2, 3, 5, 7, 11, 13, 17, 19)

	for puzzle, rounds := range roundData {
		initMonkeys()
		inspectCounts := [8]int{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		}

		for round := 0; round < rounds; round++ {
			for i, currentMonkey := range monkeys {
				for _, item := range currentMonkey.items {
					inspectedItem := currentMonkey.operation(item)
					if puzzle == 0 {
						inspectedItem = inspectedItem / 3
					} else {
						inspectedItem = inspectedItem % lcm
					}
					inspectCounts[i] += 1
					test := currentMonkey.test(inspectedItem)
					toMonkeyIndex := currentMonkey.testTrue
					if !test {
						toMonkeyIndex = currentMonkey.testFalse
					}
					monkeys[toMonkeyIndex].items = append(monkeys[toMonkeyIndex].items, inspectedItem)
				}
				currentMonkey.items = make([]int, 0)
			}
		}

		most, secondMost := 0, 0
		for _, count := range inspectCounts {
			if count > most {
				secondMost = most
				most = count
				continue
			}

			if count > secondMost {
				secondMost = count
			}
		}

		solution := most * secondMost
		roundData[puzzle] = solution
	}

	fmt.Printf("Solution is %d\n", roundData[0])
	fmt.Printf("Solution is %d\n", roundData[1])
}

func Solution() {
	both()
}
