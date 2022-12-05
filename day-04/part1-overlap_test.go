package dayfour

import (
	"testing"
)

type pointsInRucksackTest struct {
	name string
	pairs []intervalPair
	result int
}

var pointsInRucksackTests = []pointsInRucksackTest{
	{
		"One Pair - No Overlap", 
		[]intervalPair{
			{
				interval{0,1}, 
				interval{2, 3},
			},
		}, 
		0,
	},
	{
		"One Pair - Overlap", 
		[]intervalPair{
			{
				interval{0,5}, 
				interval{2, 3},
			},
		}, 
		1,
	},

	{
		"One Pair - Overlap on Edge", 
		[]intervalPair{
			{
				interval{4,6}, 
				interval{6, 6},
			},
		}, 
		1,
	},

	{
		"Two Pair - No Overlaps", 
		[]intervalPair{
			{
				interval{0,4}, 
				interval{6, 6},
			},
			{
				interval{1,3}, 
				interval{4, 5},
			},
		}, 
		0,
	},
	{
		"Two Pair - One Overlap", 
		[]intervalPair{
			{
				interval{2,3}, 
				interval{1, 5},
			},
			{
				interval{1,3}, 
				interval{4, 5},
			},
		}, 
		1,
	},
	{
		"Two Pair - Two Overlaps", 
		[]intervalPair{
			{
				interval{2,3}, 
				interval{1, 5},
			},
			{
				interval{3,3}, 
				interval{1, 9},
			},
		}, 
		2,
	},
	{
		"Combo", 
		[]intervalPair{
			{ // Yes
				interval{2,3}, 
				interval{1, 5},
			},
			{ // No
				interval{0,0}, 
				interval{1, 9},
			},
			{ // No
				interval{1,1}, 
				interval{3, 9},
			},
			{ // Yes
				interval{4,4}, 
				interval{1, 5},
			},
		}, 
		2,
	},
}


func TestOverlappingIntervals(t *testing.T) {
	for _, test := range pointsInRucksackTests {
		if output := overlappingIntervals(test.pairs); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name,  test.result, output)
		}
	}
}
