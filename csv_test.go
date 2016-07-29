package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type ToTestCsv struct {
	AnString    string  `json:"a_string"`
	AnInt       int     `json:"a_int"`
	AnInt64     int64   `json:"a_int64"`
	AnFloat32   float32 `json:"a_float32"`
	AnFloat64   float64 `json:"a_float64"`
	AnBool      bool    `json:"a_bool"`
	AnInterface interface{}
}

func TestGetHeader(t *testing.T) {

	Convey("When passed a struct", t, func() {

		headers := getHeader(ToTestCsv{})

		Convey("it should retrieve field names as comma-separated values", func() {

			So(headers, ShouldEqual, "AnString,AnInt,AnInt64,AnFloat32,AnFloat64,AnBool,AnInterface")
		})
	})
}

func TestGetValue(t *testing.T) {
	Convey("When passed a struct", t, func() {

		test := ToTestCsv{
			AnString:    "this is a string",
			AnInt:       1,
			AnInt64:     123456789,
			AnFloat32:   0.9896,
			AnFloat64:   0.58929896,
			AnBool:      true,
			AnInterface: "not nil",
		}
		values := getValues(test)

		Convey("it should retrieve values as comma-separated values", func() {

			So(values, ShouldEqual, fmt.Sprintf("this is a string,1,123456789,%.10f,%.10f,1,not nil", test.AnFloat32, test.AnFloat64))
		})
	})

	Convey("When passed a default struct", t, func() {

		test := ToTestCsv{}
		values := getValues(test)

		Convey("it should still have a comma for each field", func() {

			So(values, ShouldEqual, ",0,0,0.0000000000,0.0000000000,0,")
		})
	})
}
