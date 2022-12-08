package dayseven

import (
	"testing"
)

func TestSizeOfAllUnder(t *testing.T) {
	root := newDirectory("/")

	dirA := root.makeChildDirectory("dirA")
	dirB := root.makeChildDirectory("dirB")

	childA := dirA.makeChildDirectory("childA")
	childA.files = append(childA.files, newFile("a", 100), newFile("b", 200))

	childB := dirA.makeChildDirectory("childB")
	childB.files = append(childB.files, newFile("b", 150))
	
	childC := dirB.makeChildDirectory("childC")
	childC.files = append(childC.files, newFile("c", 100))

	target := 350
	if sizeOfAllUnder(root, 200) != target {
		t.Errorf("Incorrect total size reported, wanted %v, got %v", target, sizeOfAllUnder(root, 200))
	}
}
