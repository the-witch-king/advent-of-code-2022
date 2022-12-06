package dayfive

import (
	"testing"
)

func TestPopMany(t *testing.T) {
	stack := stack{}
	stack.pushMany([]rune{'a', 'b', 'c', 'd', 'e'})

	r := stack.popMany(3)

	if printRunes(r) != printRunes([]rune{'c', 'd', 'e'}) {
		t.Errorf("Popped incorrect values. Got %v; stack is: %v", printRunes(r), stack.print())
	}
}

func TestPushMany(t *testing.T) {
	stack := stack{}
	stack.pushMany([]rune{'a', 'b'})
	stack.pushMany([]rune{'c', 'd'})

	if stack.print() != "abcd" {
		t.Errorf("Pushed incorrect values. Got %v", stack.print())
	}
}

func printRunes(runes []rune) string {
	r := ""
	for _, c := range runes {
		r += string(c)
	}

	return r
}
