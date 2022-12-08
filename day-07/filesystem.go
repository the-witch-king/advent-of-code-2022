package dayseven

import (
	"aoc/v2/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const sizeRequiredForUpdate = 30_000_000
const totalDiskSpace = 70_000_000
func SmallestDirToDelete(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()
	os, _ := buildOperatingSystem(f)

	remainingDiskSpace := totalDiskSpace - os.Size()
	amountNeededToDelete := sizeRequiredForUpdate - remainingDiskSpace
	r := sizeOfClosest(os, amountNeededToDelete)

	return r
}

func TotalSizeOfDirectoriesToDelete(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()
	os, _ := buildOperatingSystem(f)

	return sizeOfAllUnder(os, 100000)
}

func buildOperatingSystem(input io.Reader) (*Directory, *Directory) {
	scanner := bufio.NewScanner(input)

	root := newDirectory("/")
	current := root

	for scanner.Scan() {
		cmd := scanner.Text()

		isCommand := cmd[0] == '$'
		parts := strings.Split(cmd, " ")
		if isCommand {
			isCd := parts[1] == "cd"

			if isCd {
				dest := parts[2]
				switch dest {
				case "/":
					current = root
					continue
				case "..":
					current = current.parent
					continue
				default:
					n := current.FindDir(dest)
					current = n
				}
			}
			continue
		}

		indicator, name := parts[0], parts[1]
		if indicator == "dir" {
			current.makeChildDirectory(name)
		} else {
			size, err := strconv.Atoi(indicator)

			if err != nil {
				fmt.Printf("WHOOPS! Got an invalid numbah!")
				os.Exit(69)
			}

			f := &File{}
			f.name = name
			f.size = size

			current.files = append(current.files, f)
		}
	}

	return root, current
}
