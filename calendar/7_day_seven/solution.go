package __day_seven

import (
	"embed"
	"fmt"
	"github.com/jarvainen/advent-of-code-22/calendar/util"
	"strconv"
	"strings"
)

//go:embed input.txt
var content embed.FS

type directory struct {
	directories     []*directory
	parentDirectory *directory
	files           []*file
	name            string
}

type file struct {
	name string
	size int
}

func (dir directory) findDirectory(name string) *directory {
	for _, d := range dir.directories {
		if d.name == name {
			return d
		}
	}
	return nil
}

func (dir directory) findFile(name string) *file {
	for _, f := range dir.files {
		if f.name == name {
			return f
		}
	}
	return nil
}

func (dir directory) totalSize() int {
	totalSize := 0
	for _, f := range dir.files {
		totalSize += f.size
	}
	for _, d := range dir.directories {
		totalSize += d.totalSize()
	}
	return totalSize
}

func (dir directory) getAllDirectories() []*directory {
	allDirs := []*directory{&dir}
	for _, d := range dir.directories {
		allDirs = append(allDirs, d.getAllDirectories()...)
	}
	return allDirs
}

func processFileStructure() directory {
	var fileStructure = directory{name: "/"}
	currentDirectory := &fileStructure
	isListing := false
	util.ReadFileByLine(content, "input.txt", func(line string) {
		fields := strings.Fields(line)
		if isListing && fields[0] != "$" {
			switch fields[0] {
			case "dir":
				name := fields[1]
				if currentDirectory.findDirectory(name) == nil {
					newDir := directory{
						parentDirectory: currentDirectory,
						name:            fields[1],
					}
					currentDirectory.directories = append(currentDirectory.directories, &newDir)
				}
				return
			default:
				fileSize, _ := strconv.Atoi(fields[0])
				name := fields[1]
				if currentDirectory.findFile(name) == nil {
					newFile := file{
						name: name,
						size: fileSize,
					}
					currentDirectory.files = append(currentDirectory.files, &newFile)
				}
				return
			}
		}

		if fields[0] == "$" {
			isListing = false
			switch fields[1] {
			case "cd":
				name := fields[2]
				switch name {
				case "/":
					return
				case "..":
					currentDirectory = currentDirectory.parentDirectory
				default:
					currentDirectory = currentDirectory.findDirectory(name)
				}
				break
			case "ls":
				isListing = true
				break
			}
		}
	})
	return fileStructure
}

func silver() {
	fileStructure := processFileStructure()
	allDirs := fileStructure.getAllDirectories()
	solution := 0
	for _, d := range allDirs {
		totalSize := d.totalSize()
		if totalSize < 100000 {
			solution += totalSize
		}
	}
	fmt.Printf("Solution is %d\n", solution)
}

func gold() {
	fileStructure := processFileStructure()
	fileStructureSize := fileStructure.totalSize()
	allDirs := fileStructure.getAllDirectories()
	limit := fileStructureSize - 40000000
	solution := fileStructureSize
	for _, d := range allDirs {
		totalSize := d.totalSize()
		if limit < totalSize && solution > totalSize {
			solution = totalSize
		}
	}
	fmt.Printf("Solution is %d\n", solution)
}

func Solution() {
	silver()
	gold()
}
