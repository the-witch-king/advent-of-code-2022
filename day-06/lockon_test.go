package daysix

import (
	"testing"
)

type startOfTest struct {
	name   string
	s      string
	result int
}

var startOfPacketTests = []startOfTest{
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
		"abcazcd",
		"abcazcd",
		4,
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

	// Provided
	{
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		5 - 1,
	},
	{
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nppdvjthqldpwncqszvftbrmjlhg",
		6 - 1,
	},
	{
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		10 - 1,
	},
	{
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		11 - 1,
	},
}

func TestStartOfPacket_length4(t *testing.T) {
	for _, test := range startOfPacketTests {
		if output := indexOfUnique(test.s, 4); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name, test.result, output)
		}
	}
}

var startOfMessageTests = []startOfTest{
	// Provided
	{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		19 - 1,
	},
	{
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		23 - 1,
	},
	{
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nppdvjthqldpwncqszvftbrmjlhg",
		23 - 1,
	},
	{
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		29 - 1,
	},
	{
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		26 - 1,
	},
}

func TestStartOfMessage_length14(t *testing.T) {
	for _, test := range startOfMessageTests {
		if output := indexOfUnique(test.s, 14); output != test.result {
			t.Errorf("\n[%s]\nWanted:\t%v\nGot:\t%v", test.name, test.result, output)
		}
	}
}
