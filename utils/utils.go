package utils

import (
	"fmt"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Unable to read input!", err)
		os.Exit(69)
	}

  return file
}
