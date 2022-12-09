package dayeight

import "testing"

type scenicScoresTest struct {
	name   string
	input  [][]int
	result int
}

var scenicScoreTests = []scenicScoresTest{
	{
		"foo",
		[][]int{
			{0, 0, 0},
			{0, 5, 0},
			{0, 0, 0},
		},
		1,
	},
	{
		"Smol",
		[][]int{
			{0, 1},
			{2, 3},
		},
		0,
	},
	{
		"Bigger",
		[][]int{
			{0, 1, 0},
			{0, 2, 0},
			{0, 1, 0},
		},
		1,
	},
	{
		"Biggest",
		[][]int{
			{0, 1, 2, 0},
			{0, 2, 7, 3},
			{0, 1, 4, 5},
			{5, 1, 7, 3},
		},
		4,
	},
	{
		"Provided",
		[][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
		8,
	},
}

func TestScenicScores(t *testing.T) {
	for _, test := range scenicScoreTests {
		if output := scenicScore(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name, test.result, output)
		}
	}
}
