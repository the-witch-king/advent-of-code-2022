package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
  Calories int
}

type Elves []Elf

func (e Elves) Len() int {
  return len(e)
}
func (e Elves) Less(a, b int) bool {
  return e[a].Calories < e[b].Calories
}
func (e Elves) Swap(a, b int) {
  e[a], e[b] = e[b], e[a]
}



func main() {
  file, err := os.Open("./day-01/input.txt")

  if err != nil {
    fmt.Println("Unable to read input!", err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  elves := Elves{}

  elf := &Elf{}

  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      elves = append(elves, *elf) 
      elf = &Elf{}
      continue
    }

    cals, err := strconv.Atoi(line)

    if err != nil {
      fmt.Println("Error reading a line!")
      os.Exit(69)
    }

    elf.Calories += cals
  }

  /* This was Part 1 */
  // maxElf, err := MostCalories(elves)

  // if err != nil {
  //   fmt.Println("Didn't get no elves, whooooops")
  //   os.Exit(42)
  // }

  // fmt.Println("Most calories carried are: ", maxElf.Calories)

  sort.Sort(Elves(elves))

  totalCalories := 0
  for _, e := range elves[len(elves) -3:] {
    totalCalories += e.Calories
  }

  fmt.Println("Calories carried by top three elves are: ", totalCalories)
}

func MostCalories(elves []Elf) (Elf, error) {
  if len(elves) == 0 {
    var empty Elf
    return empty, errors.New("Empty list provided.")
  }

  max := elves[0]

  for _, e := range elves {
    if e.Calories >= max.Calories {
      max = e
    }
  }

  return max, nil
}
