package daytwo_part1

import (
	"aoc/v2/utils"
	"bufio"
	"io"
	"strings"
)

func TotalScore(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()
	return totalScore(f)
}

const drawPoints = 3
const winPoints = 6
const lossPoints = 0

var choicePoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}
var moves = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

func totalScore(data io.Reader) int {
	scanner := bufio.NewScanner(data)

	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		moves := strings.Split(line, " ")

		opp, my := moves[0], moves[1]

		score += game(opp, my)
	}

	return score
}

// Takes two moves, returns outcome
func game(opp, mine string) int {
	cp := choicePoints[mine]

	oppMove := moves[opp]

	// Draw
	if oppMove == mine {
		return drawPoints + cp
	}

	// Win
	if isWin(oppMove, mine) {
		return winPoints + cp
	}

	// Loss
	return lossPoints + cp
}

func isWin(a, b string) bool {
	return a == "X" && b == "Y" || a == "Y" && b == "Z" || a == "Z" && b == "X"
}
