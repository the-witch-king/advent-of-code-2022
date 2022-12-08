package dayseven

import (
	"testing"
)

func TestDirectorySize(t *testing.T) {
	dirB := newDirectory("dirB")
	dirB.files = append(dirB.files, newFile("c", 100))

	dirA := dirB.makeChildDirectory("dirA")
	dirA.files = append(dirA.files, newFile("a", 50), newFile("b", 75))

	// Check dirA working
	if dirA.Size() != 125 {
		t.Errorf("DirA Size is incorrect. Got %v, wanted %v", dirA.Size(), 125)
	}

	// Check dirB working
	if dirB.Size() != 225 {
		t.Errorf("DirB Size is incorrect. Got %v, wanted %v", dirA.Size(), 225)
	}
}
