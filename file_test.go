package main

import (
	"fmt"
	"os"
	"path"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadFromOutputAndWriteToJson(t *testing.T) {
	file := "inventory.json"
	bytes, _ := readFile("output.log")
	writeFile(file, bytes)
}

func TestCreateWriteReadFile(t *testing.T) {
	file := "file_test_TestCreateWriteReadFile.txt"
	fullpath := path.Join(pwd, file)
	fmt.Printf("fullpath: %s\n", fullpath)

	Convey("When writing to the file", t, func() {

		text := "this is just some example text to test with"
		writeFile(file, []byte(text))

		Convey("it should write the correct text", func() {
			bytes, readFileErr := readFile(file)
			So(readFileErr, ShouldBeNil)
			So(string(bytes), ShouldEqual, text)
		})

		Reset(func() {
			os.Remove(fullpath)
		})
	})
}
