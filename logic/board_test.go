package logic

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {

	Convey("Given some board", t, func() {
		board := FromString("__X _X_", 3, 2, 0)
		Convey("Is valid", func() {
			So(board.width, ShouldEqual, 3)
			So(board.height, ShouldEqual, 2)
			So(board.cells, ShouldResemble, []bool{false, false, true, false, true, false})
		})
		Convey("Get", func() {
			So(board.Get(0, 0), ShouldEqual, false)
			So(board.Get(1, 0), ShouldEqual, false)
			So(board.Get(2, 0), ShouldEqual, true)
			So(board.Get(0, 1), ShouldEqual, false)
			So(board.Get(1, 1), ShouldEqual, true)
			So(board.Get(2, 1), ShouldEqual, false)
		})
		Convey("Set", func() {
			So(board.Get(2, 1), ShouldEqual, false)
			So(board.Set(2, 1, true), ShouldEqual, false)
			So(board.Get(2, 1), ShouldEqual, true)
			So(board.Set(2, 1, false), ShouldEqual, true)
		})
		Convey("ToString", func() {
			So(board.ToString(), ShouldEqual, "__X _X_")
		})
		Convey("livingNeighbours", func() {
			So(board.livingNeighbours(0, 0), ShouldEqual, 1)
			So(board.livingNeighbours(1, 0), ShouldEqual, 2)
			So(board.livingNeighbours(2, 0), ShouldEqual, 1)
		})
	})
	Convey("Given a static object", t, func() {
		board := FromString("_X_ X_X _X_", 3, 3, 0)
		So(board.ToString(), ShouldEqual, "_X_ X_X _X_")
		board.NextGen()
		So(board.ToString(), ShouldEqual, "_X_ X_X _X_")
	})
	Convey("Given a blinker object", t, func() {
		board := FromString("___ XXX ___", 3, 3, 0)
		So(board.ToString(), ShouldEqual, "___ XXX ___")
		board.NextGen()
		So(board.ToString(), ShouldEqual, "_X_ _X_ _X_")
		board.NextGen()
		So(board.ToString(), ShouldEqual, "___ XXX ___")
	})

	Convey("Given a different board", t, func() {
		board := FromString("X_ XX _X", 2, 3, 0)
		So(board.ToString(), ShouldEqual, "X_ XX _X")
		So(board.Texture(), ShouldResemble, []byte{
			0, 255, 0, 0,
			0, 255, 255, 0,
			0, 0, 255, 0,
			0, 0, 0, 0})
	})

	Convey("A board with offset", t, func() {
		board := FromString("X_ XX _X", 2, 3, 1)
		So(board.ToString(), ShouldEqual, "____ _X__ _XX_ __X_ ____")
	})
}
