package main

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestFileCreation(t *testing.T) {

	// Arrange
	file := "test.txt"
	fullpath := path.Join(pwd, file)
	fmt.Println(fullpath)

	// Act
	recreateFile(file)

	// Assert
	_, err := os.Stat(fullpath)
	if os.IsNotExist(err) {
		t.Errorf("Expecting file '%s' to exist", fullpath)
	}

}
