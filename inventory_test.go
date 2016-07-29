package main

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInventorySerialization(t *testing.T) {

	Convey("When unmarshalling inventory", t, func() {

		bytes, _ := readFile("inventory_test.json")

		fmt.Printf("\n%s\n...\n", string(bytes[:64]))

		var inventory inventory
		unmarshalError := json.Unmarshal(bytes, &inventory)

		Convey("it should succeed", func() {
			So(unmarshalError, ShouldBeNil)
			So(inventory.Success, ShouldBeTrue)
			So(inventory.InventoryDelta.NewTimestampMS, ShouldEqual, 1469727262791)
			So(inventory.InventoryDelta.InventoryItems, ShouldHaveLength, 261)
		})
	})

}
