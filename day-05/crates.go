package dayfive

import (
	"aoc/v2/utils"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const spacesPerMapToken = 4
const endWhitespaceOffset = 1 // Map doesn't have trailling whitespace at end of it
const crateLabelIndex = 1

// Indices for commands
const amountIndex = 1
const fromIndex = 3
const toIndex = 5

func MoveCrates(path string) string {
	f := utils.OpenFile(path)
	defer f.Close()

	return moveCrates(f, crateMover9001)
}

// Read file, split into "map" + commands
func moveCrates(data io.Reader, crane func([]stack, string)) string {
	scanner := bufio.NewScanner(data)

	// Parse map into stacks
	stacks := parseMapIntoStacks(scanner)


	// Parse instructions, perform movements
	for scanner.Scan() {
		line := scanner.Text()
		crane(stacks, line)
	}

	result := ""

	for _, stack := range stacks {
		result += string(stack.peek())
	}

	return result
}

// Part 2
func crateMover9001(stacks []stack, commandLine string) {
		cmd := strings.Split(commandLine, " ")

		amount, err := strconv.Atoi(cmd[amountIndex])
		handleError(err)
		from, err := strconv.Atoi(cmd[fromIndex])
		handleError(err)
		to, err := strconv.Atoi(cmd[toIndex])
		handleError(err)

		c := stacks[from - 1].popMany(amount)
		stacks[to - 1].pushMany(c)
}


// Part 1
func crateMover9000(stacks []stack, commandLine string) {
		cmd := strings.Split(commandLine, " ")

		amount, err := strconv.Atoi(cmd[amountIndex])
		handleError(err)
		from, err := strconv.Atoi(cmd[fromIndex])
		handleError(err)
		to, err := strconv.Atoi(cmd[toIndex])
		handleError(err)


		for i := 0; i < amount; i++ {
			c := stacks[from - 1].pop()
			stacks[to - 1].push(c)
		}
}

// Parse map into stacks
func parseMapIntoStacks(scanner *bufio.Scanner) []stack {
	var stacks []stack
	for scanner.Scan() {
		line := scanner.Text()

		// Done with map
		if line == "" {
			break
		}

		// line 9, indicating count of stack
		if line[0] == ' ' {
			continue
		}

		// First pass
		var amountOfStacks int
		if len(stacks) == 0 {
			amountOfStacks = (len(line) + endWhitespaceOffset) / spacesPerMapToken
			stacks = make([]stack, amountOfStacks)
		}

		slices := chunkSlice([]rune(line), spacesPerMapToken)

		for i := 0; i < len(slices); i++ {
			c := slices[i][crateLabelIndex]


			if c == ' ' {
				continue
			}

			stacks[i].unshift(c)
		}
	}

	return stacks
}

func chunkSlice[T any](slice []T, chunkSize int) [][]T {
	var results [][]T
	size := int(math.Ceil(float64(len(slice)) / float64(chunkSize)))
	results = make([][]T, size)

	index := 0
	for cursor := 0; cursor < len(slice); cursor += chunkSize {
		if cursor+chunkSize > len(slice) {
			results[index] = slice[cursor:]
		} else {
			results[index] = slice[cursor : cursor+chunkSize]
		}
		index++
	}

	return results
}

func handleError(err error) {
	if err == nil {
		return
	}

	fmt.Println("Error: ", err)
	os.Exit(69)
}
