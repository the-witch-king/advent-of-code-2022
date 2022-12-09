package dayeight

import "testing"

type visibleTreesTest struct {
	name   string
	input  [][]int
	result int
}

var visibleTreesTests = []scenicScoresTest{
	{
		"Smol",
		[][]int{
			{0, 1},
			{2, 3},
		},
		4,
	},
	{
		"Bigger",
		[][]int{
			{0, 1, 0},
			{0, 2, 0},
			{0, 1, 0},
		},
		9,
	},
	{
		"Biggest",
		[][]int{
			{0, 1, 2, 0},
			{0, 2, 7, 3},
			{0, 1, 4, 5},
			{5, 1, 7, 3},
		},
		16,
	},
	{
		"Hidden",
		[][]int{
			{4, 5, 4},
			{5, 2, 5},
			{4, 5, 4},
		},
		8,
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
		21,
	},
}

func TestVisibleTrees(t *testing.T) {
	for _, test := range visibleTreesTests {
		if output := visibleTrees(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name, test.result, output)
		}
	}
}

