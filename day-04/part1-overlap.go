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
func (a interval) overlaps(b interval) bool {
	return (a.Start <= b.Start && b.End <= a.End) ||
		(b.Start <= a.Start && a.End <= b.End)
}

type intervalPair struct {
	First  interval
	Second interval
}

func overlappingIntervals(intervals []intervalPair) int {
	total := 0

	for _, pair := range intervals {
		a, b := pair.First, pair.Second

		if a.overlaps(b) {
			total++
		}
	}

	return total
}

/** Input parsing **/
func Overlap(path string) int {
	f := utils.OpenFile(path)
	defer f.Close()

	return overlappingIntervals(getIntervals(f))
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


