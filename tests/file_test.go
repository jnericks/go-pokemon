package test

import (
	"fmt"
	"os"
	"path"
	"testing"

	. "github.com/jnericks/go-pokemon"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateWriteReadFile(t *testing.T) {
	file := "file_test_TestCreateWriteReadFile.txt"
	fullpath := path.Join(GetDataPath(), file)
	fmt.Printf("fullpath: %s\n", fullpath)

	Convey("When writing to the file", t, func() {

		text := "this is just some example text to test with"
		WriteFile(file, []byte(text))

		Convey("it should write the correct text", func() {
			bytes, readFileErr := ReadFile(file)
			So(readFileErr, ShouldBeNil)
			So(string(bytes), ShouldEqual, text)
		})

		Convey("and we append to the file it should append the correct text", func() {

			append := "append me!"
			AppendFile(file, append)

			bytes, readFileErr := ReadFile(file)
			So(readFileErr, ShouldBeNil)
			So(string(bytes), ShouldEqual, fmt.Sprintf("%s\n%s", text, append))
		})

		Reset(func() {
			os.Remove(fullpath)
		})
	})
}
