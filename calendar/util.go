package calendar

import (
	"bufio"
	"log"
	"os"
)

type Solution func()

func readFileByLine(fileName string, processLine func(string)) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		processLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
