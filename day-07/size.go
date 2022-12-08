package dayseven

import "math"


func sizeOfClosest(node *Directory, minimum int) int {
	return soc(node, minimum, math.MaxInt)
}
func soc(node *Directory, minimum int, lowest int) int {
	nodeSize := node.Size()
	if nodeSize > minimum {
		lowest = min(nodeSize, lowest)
	}

	for _, d := range node.directories {
		size := soc(d, minimum, lowest)
		lowest = min(lowest, size)
	}

	return lowest
}

func sizeOfAllUnder(root *Directory, max int) int {
	s := 0

	if root == nil {
		return 0
	}

	if root.Size() < max {
		s += root.Size()
	}

	for _, d := range root.directories {
		s += sizeOfAllUnder(d, max)
	}

	return s
}


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
