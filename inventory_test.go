package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInventorySerialization(t *testing.T) {

	Convey("When adding 1 + 4 it should equal 5", func() {
		So(1+4, ShouldEqual, 5)
	})
}
