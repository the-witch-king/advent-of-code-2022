package daysix

import (
	"aoc/v2/utils"
	"bufio"
)

const oneIndexOffset = 1
const packetLen = 4
const messageLen = 14

func LockOnPacket(path string) int {
  f := utils.OpenFile(path)
  defer f.Close()

  scanner := bufio.NewScanner(f)
  scanner.Scan()
  line := scanner.Text()
  return indexOfUnique(line, packetLen) + oneIndexOffset
}

func LockOnMessage(path string) int {
  f := utils.OpenFile(path)
  defer f.Close()

  scanner := bufio.NewScanner(f)
  scanner.Scan()
  line := scanner.Text()
  return indexOfUnique(line, messageLen) + oneIndexOffset
}

func indexOfUnique(s string, maxCount int) int {
  v := map[rune]int{}
  count := 0

  for i := 0; i < len(s); i++ {
    c := rune(s[i])

    if lastSeenAt, visited := v[c]; visited {
      count = 0
      i = lastSeenAt
      v = map[rune]int{}
      continue
    }

    v[c] = i
    count++

    if count == maxCount {
      return i
    }
  }
  return -1
}

