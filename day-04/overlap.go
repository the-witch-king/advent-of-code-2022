package dayfour

import (
	"aoc/v2/utils"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/** Interval **/
type interval struct {
	Start int
	End   int
}
func overlapCompletely(a, b interval) bool {
	return (a.Start <= b.Start && b.End <= a.End) ||
		(b.Start <= a.Start && a.End <= b.End)
}
func overlap(a, b interval) bool {
	// A
	// Start within range
	aS := a.Start >= b.Start && a.Start <= b.End

	// End within range
	aE := a.End >= b.Start && a.End <= b.End

	// B
	// Start within range
	bS := b.Start >= a.Start && b.Start <= a.End

	// End within range
	bE := b.End >= a.Start && b.End <= a.End

	return aS || aE || bS || bE
}

type intervalPair struct {
	First  interval
	Second interval
}

func overlappingIntervals(intervals []intervalPair, overlapCheck func (a, b interval) bool) int {
	total := 0

	for _, pair := range intervals {
		a, b := pair.First, pair.Second

		if overlapCheck(a, b) {
			total++
		}
	}

	return total
}

/** Input parsing **/
func OverlapCompletely(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()

	return overlappingIntervals(getIntervals(f), overlapCompletely)
}
func Overlap(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()

	return overlappingIntervals(getIntervals(f), overlap)
}

func getIntervals(d io.Reader) []intervalPair {
	scanner := bufio.NewScanner(d)

	pairs := []intervalPair{}

	for scanner.Scan() {
		line := scanner.Text()
		pairs = append(pairs, intervalsFromLine(line))
	}

	return pairs
}

func intervalsFromLine(l string) intervalPair {
	p := strings.Split(l, ",")
	first, second := p[0], p[1]

	return intervalPair{intervalFromString(first), intervalFromString(second)}
}

func intervalFromString(s string) interval {
	pair := strings.Split(s, "-")
	firstStart, firstEnd := pair[0], pair[1]

	start, err := strconv.Atoi(firstStart)
	if err != nil {
		fmt.Println("Couldn't parse input")
		os.Exit(1)
	}

	end, err := strconv.Atoi(firstEnd)
	if err != nil {
		fmt.Println("Couldn't parse input")
		os.Exit(1)
	}

	return interval{start, end}
}


