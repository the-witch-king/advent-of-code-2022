package daythree_part1

import (
	"io"
	"strings"
	"testing"
)

type pointsInRucksackTest struct {
	name string
	input io.Reader
	result int
}

var pointsInRucksackTests = []pointsInRucksackTest{
	// Simple
	{"a", strings.NewReader("afaJ"), 1},
	{"A", strings.NewReader("AfAJ"), 27},
	
	// Provided
	{"p", strings.NewReader("vJrwpWtwJgWrhcsFMMfFFhFp"), 16},
	{"L", strings.NewReader("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"), 38},
	{"P", strings.NewReader("PmmdzqPrVvPwwTWBwg"), 42}, 

	// Combo
	{"Combo", strings.NewReader("afaJ\nAfAJ\nvJrwpWtwJgWrhcsFMMfFFhFp"), 16 + 1 + 27},
}

func TestTotalScore(t *testing.T) {
	for _, test := range pointsInRucksackTests {
		if output := pointsInRucksack(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name,  test.result, output)
		}
	}
}
