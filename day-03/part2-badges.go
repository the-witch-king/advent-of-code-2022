package daythree

import (
	"aoc/v2/utils"
	"bufio"
	"io"
)

const groupCount = 3

func SumOfBadges(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()
	return sumOfBadges(f)
}

func sumOfBadges(data io.Reader) int {
	scanner := bufio.NewScanner(data)

	score := 0
	groupCounter := 0
	groupMap := map[rune]int{}

	for scanner.Scan() {
		line := scanner.Text()
		groupCounter++

		for _, c := range line {
			if groupCounter == 1 {
				groupMap[c] = groupCounter
			} else {
				if groupMap[c] == groupCounter-1 {
					groupMap[c] = groupCounter
				}
			}
		}


		// Next group, reset
		if groupCounter == 3 {
			for c, v := range groupMap {
				if v == 3 {
					score += pointsForChar(c)
					break
				}
			}

			groupCounter = 0
			groupMap = map[rune]int{}
		}
	}

	return score
}
