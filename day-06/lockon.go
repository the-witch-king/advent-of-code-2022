package daysix

import (
	"aoc/v2/utils"
	"bufio"
	"io"
)

const oneIndexOffset = 1

func LockOn(path string) int {
  f := utils.OpenFile(path)
  defer f.Close()

  return lockOn(f)
}

func lockOn(data io.Reader) int {
  scanner := bufio.NewScanner(data)
  scanner.Scan()
  line := scanner.Text()
  return startOfPacket(line) + oneIndexOffset
}

func startOfPacket(s string) int {
  v := map[rune]bool{}
  count := 0
  index := 0

  for i, c := range s {
    if _, visited := v[c]; visited {
      count = 0
      v = map[rune]bool{}
      continue
    }

    v[c] = true
    index = i
    count++

    if count == 4 {
      break
    }
  }
  return index
}

