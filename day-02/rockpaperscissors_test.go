package rps

import (
	"io"
	"strings"
	"testing"
)

// A, X   Rock     1 pts
// B, Y   Paper    2 pts
// C, Z   Scissors 3 pts
// Win   6 pts
// Draw  3 pts
// Lose  0 pts

type totalScoreTest struct {
	name string
	input io.Reader
	result int
}

var totalScoreTests = []totalScoreTest{
	// Rock
	{"R - D", strings.NewReader("A X\n"), 4}, // D
	{"R - L", strings.NewReader("B X\n"), 1}, // L
	{"R - W", strings.NewReader("C X\n"), 7}, // W

	// Paper
	{"P - W", strings.NewReader("A Y\n"), 8}, // W
	{"P - D", strings.NewReader("B Y\n"), 5}, // D
	{"P - L", strings.NewReader("C Y\n"), 2}, // L

	// Scissors
	{"S - L", strings.NewReader("A Z\n"), 3}, // L
	{"S - W", strings.NewReader("B Z\n"), 9}, // W
	{"S - D", strings.NewReader("C Z\n"), 6}, // D

	// Combos

	// Draws
	{"C - D", strings.NewReader("C Z\nB Y\nA X\n"), 15},
	// Wins
	{"C - W", strings.NewReader("C X\nA Y\nB Z\n"), 24},
	// Losses
	{"C - L", strings.NewReader("B X\nC Y\nA Z\n"), 6},
	// Mix
	{"C - M", strings.NewReader("C X\nB Y\nA Z\n"), 15},
}

func TestTotalScore(t *testing.T) {
	for _, test := range totalScoreTests {
		if output := totalScore(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name,  test.result, output)
		}
	}
}
