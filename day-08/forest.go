package dayeight

import (
	"aoc/v2/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func VisibleTrees(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()

	return visibleTrees(makeMatrix(f))
}

func makeMatrix(data io.Reader) [][]int {
	matrix := [][]int{}

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := scanner.Text()

		row := []int{}
		for _, c := range line {
			num, err := strconv.Atoi(string(c))

			if err != nil {
				fmt.Println("Unable to parse input, invalid char found")
				os.Exit(69)
			}

			row = append(row, num)
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func visibleTrees(matrix [][]int) int {
	count := 0
	max := -1

	rows := len(matrix)
	cols := len(matrix[0])

	seen := make([][]bool, rows)
	for i := 0; i < len(matrix); i++ {
		seen[i] = make([]bool, cols)
	}

	// LtR
	for r := 0; r < rows; r++ {
		row := matrix[r]

		// Left
		max = -1
		for c := 0; c < cols; c++ {
			val := row[c]
			if val > max {
				if !seen[r][c] {
					count++
					seen[r][c] = true
				}
				max = val
			}

		}

		// Right
		max = -1
		for c := cols - 1; c > -1; c-- {
			val := row[c]
			if val > max {
				if !seen[r][c] {
					count++
					seen[r][c] = true
				}
				max = val
			}

		}
	}

	// TtB
	for c := 0; c < cols; c++ {
		// Top
		max = -1
		for r := 0; r < rows; r++ {
			val := matrix[r][c]
			if val > max {
				if !seen[r][c] {
					count++
					seen[r][c] = true
				}
				max = val
			}
		}

		// Bottom
		max = -1
		for r := rows - 1; r > -1; r-- {
			val := matrix[r][c]
			if val > max {
				if !seen[r][c] {
					count++
					seen[r][c] = true
				}
				max = val
			}
		}
	}

	return count
}
