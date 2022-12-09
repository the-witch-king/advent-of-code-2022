package dayeight

import (
	"reflect"
	"strings"
	"testing"
)

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
