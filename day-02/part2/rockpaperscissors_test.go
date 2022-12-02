package daytwo_part2

import (
	"io"
	"strings"
	"testing"
)

// A	Rock			1 pts
// B  Paper			2 pts
// C  Scissors	3 pts

// X	Lose			0 pts
// Y	Draw			3 pts
// Z	Win				6 pts

type totalScoreTest struct {
	name   string
	input  io.Reader
	result int
}

var totalScoreTests = []totalScoreTest{
	// Opp Rock
	{"R - L", strings.NewReader("A X\n"), 3}, // L, S
	{"R - D", strings.NewReader("A Y\n"), 4}, // D, R
	{"R - W", strings.NewReader("A Z\n"), 8}, // W, P

	// Opp Paper
	{"P - L", strings.NewReader("B X\n"), 1}, // L, R
	{"P - D", strings.NewReader("B Y\n"), 5}, // D, P
	{"P - W", strings.NewReader("B Z\n"), 9}, // W, S

	// Opp Scissors
	{"S - L", strings.NewReader("C X\n"), 2}, // L, P
	{"S - D", strings.NewReader("C Y\n"), 6}, // D, S
	{"S - W", strings.NewReader("C Z\n"), 7}, // W, R

	// Combos

	// Draws
	{"C - D", strings.NewReader("A Y\nB Y\nC Y\n"), 4 + 5 + 6},
	// Wins
	{"C - W", strings.NewReader("A Z\nB Z\nC Z\n"), 8 + 9 + 7},
	// Losses
	{"C - L", strings.NewReader("A X\nB X\nC X\n"), 3 + 1 + 2},
	// Mix
	{"C - M", strings.NewReader("A X\nB Y\nC Z\n"), 3 + 5 + 7},
}

func TestTotalScorePart2(t *testing.T) {
	for _, test := range totalScoreTests {
		if output := totalScore(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name, test.result, output)
		}
	}
}
