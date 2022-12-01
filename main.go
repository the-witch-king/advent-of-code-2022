package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
  "errors"
)

type Elf struct {
  Calories int
}

func main() {
  file, err := os.Open("./day-01/input.txt")

  if err != nil {
    fmt.Println("Unable to read input!", err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  elves := []Elf{}

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

  maxElf, err := MostCalories(elves)

  if err != nil {
    fmt.Println("Didn't get no elves, whooooops")
    os.Exit(42)
  }

  fmt.Println("Most calories carried are: ", maxElf.Calories)
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
