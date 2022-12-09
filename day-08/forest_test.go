package dayeight

import (
	"reflect"
	"strings"
	"testing"
)

type visibleTreesTest struct {
	name   string
	input  [][]int
	result int
}

var visibleTreesTests = []visibleTreesTest{
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

func TestCompletelyOverlappingIntervals(t *testing.T) {
	for _, test := range visibleTreesTests {
		if output := visibleTrees(test.input); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name, test.result, output)
		}
	}
}

// Test string parser
func TestMakeMatrixFromInput(t *testing.T) {
	input := strings.NewReader("04540\n22222\n12321")

	expected := [][]int{
		{0, 4, 5, 4, 0},
		{2, 2, 2, 2, 2},
		{1, 2, 3, 2, 1},
	}

	result := makeMatrix(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Did not get desired Matrix. Got %v", result)
	}
}
