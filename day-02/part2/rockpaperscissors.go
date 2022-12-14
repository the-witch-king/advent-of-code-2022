package daytwo_part2

import (
	"aoc/v2/utils"
	"bufio"
	"io"
	"strings"
)

const drawPoints = 3
const winPoints = 6
const lossPoints = 0

const LOSS = "X"
const DRAW = "Y"
const WIN = "Z"

var choicePoints = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}

var winMoves = map[string]string{
	"A": "B",
	"B": "C",
	"C": "A",
}

var lossMoves = map[string]string{
	"A": "C",
	"B": "A",
	"C": "B",
}


func TotalScore(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()
	return totalScore(f)
}

func totalScore(data io.Reader) int {
	scanner := bufio.NewScanner(data)

	score := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		moves := strings.Split(line, " ")

		opp, my := moves[0], moves[1]

		score += game(opp, my)
	}

	return score
}

// Takes two moves, returns outcome
func game(opp, outcome string) int {
	switch outcome {
	case LOSS:
		return lossPoints + choicePoints[lossMoves[opp]] 
	case DRAW:
		return drawPoints + choicePoints[opp]
	case WIN:
		return winPoints + choicePoints[winMoves[opp]]
	default:
		return 0
	}
}
