package tests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Multiply multiplies two integers and returns the result.
func Multiply(a, b int) int {
	return a * b
}

// Divide divides two integers and returns the result.
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func TestAdvancedOperations(t *testing.T) {
	var a, b int

	Convey("Given two numbers", t, func() {
		a = 6
		b = 2

		Convey("When we multiply them", func() {
			result := Multiply(a, b)
			Convey("Then the result should be correct", func() {
				So(result, ShouldEqual, 12)
			})
		})

		Convey("When we divide the first by the second", func() {
			result, err := Divide(a, b)
			Convey("Then there should be no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the result should be correct", func() {
				So(result, ShouldEqual, 3)
			})
		})

		Convey("When we divide the first by zero", func() {
			result, err := Divide(a, 0)
			Convey("Then there should be an error", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("And the result should be zero", func() {
				So(result, ShouldEqual, 0)
			})
		})
	})

	Convey("Given a negative number", t, func() {
		a = -4
		b = 2

		Convey("When we multiply them", func() {
			result := Multiply(a, b)
			Convey("Then the result should be correct", func() {
				So(result, ShouldEqual, -8)
			})
		})

		Convey("When we divide the first by the second", func() {
			result, err := Divide(a, b)
			Convey("Then there should be no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the result should be correct", func() {
				So(result, ShouldEqual, -2)
			})
		})
	})

	Convey("Given a large number", t, func() {
		a = 1e9
		b = 1e9

		Convey("When we multiply them", func() {
			result := Multiply(a, b)
			Convey("Then the result should be correct", func() {
				So(result, ShouldEqual, 1e18)
			})
		})

		Convey("When we divide the first by the second", func() {
			result, err := Divide(a, b)
			Convey("Then there should be no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And the result should be correct", func() {
				So(result, ShouldEqual, 1)
			})
		})
	})
}
