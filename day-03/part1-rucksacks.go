package daythree

import (
	"bufio"
	"io"
  "aoc/v2/utils"
)

func PointsInRucksack(path string) int {
  f := utils.OpenFile(path)
  defer f.Close()
	return pointsInRucksack(f)
}

func pointsInRucksack(data io.Reader) int {
  scanner := bufio.NewScanner(data)

  points := 0
  for scanner.Scan() {
    line := scanner.Text()
    length := len(line)
    firstHalf := line[:length/2]
    secondHalf := line[length/2:]
    m := map[rune]bool{}

    for _, c := range firstHalf {
      m[c] = true
    }

    for _, c := range secondHalf {
      if m[c] {
        points += pointsForChar(c)
        break
      }
    }
  }

  return points 
}



