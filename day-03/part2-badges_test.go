package daythree

import (
	"io"
	"strings"
	"testing"
)

type sumOfBadgesTest struct {
	name string
	input io.Reader
	result int
}

var sumOfBadgesTests = []sumOfBadgesTest{
	// Provided
	{
		"r", 
		strings.NewReader("vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\n"), 
		18,
	},
	{
		"Z", 
		strings.NewReader("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw\n"), 
		52,
	},

	// Combo
	{
		"r & Z", 
		strings.NewReader("vJrwpWtwJgWrhcFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n ttgJtRGJQctTZtZT\n CrZsJsPPZsGzwwsLwLmpwMDw\n"), 
		18 + 52,
	},
}

func TestSumOfBadges(t *testing.T) {
	for _, test := range sumOfBadgesTests {
		if output := sumOfBadges(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name,  test.result, output)
		}
	}
}
