package util

import (
	"bufio"
	"embed"
	"log"
)

func ReadFileByLine(content embed.FS, fileName string, processLine func(string)) {
	file, err := content.Open(fileName)

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
