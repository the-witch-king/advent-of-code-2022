package daysix

import (
	"testing"
)

type startOfPacketTest struct {
	name string
	s string
	result int
}

var startOfPacketTests = []startOfPacketTest{
	{
		"abcd",
		"abcd", 
		3,
	},
	{
		"abcde", 
		"abcde", 
		3,
	},
	{
		"abcabcd", 
		"abcabcd", 
		6,
	},
	{
		"abcaz", 
		"abcaz", 
		4,
	},

	{
		"abcdefghijklmno", 
		"abcdefghijklmno", 
		3,
	},

	{
		"aaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbdaaefg", 
		"aaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbdaaefg", 
		39,
	},
}

func TestOverlappingIntervals(t *testing.T) {
	for _, test := range startOfPacketTests {
		if output := startOfPacket(test.s); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name,  test.result, output)
		}
	}
}
