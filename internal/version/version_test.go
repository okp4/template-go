package version

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewInfo(t *testing.T) {
	Convey("When summoning a current info object", t, func() {
		result := NewInfo()
		Convey("Then an info object is returned", func() {
			So(result.GoVersion, ShouldStartWith, "go version")
		})
	})
}

func TestInfoString(t *testing.T) {
	Convey("Given an info object", t, func() {
		info := new(Info)
		info.Name = "go-template"
		info.Version = "0.0"
		info.GitCommit = "13245"
		info.GoVersion = "go fake version 42"

		Convey("When the object string function is called", func() {
			result := info.String()
			Convey("Then the result should be an info string", func() {
				expected := `go-template: 0.0
git commit: 13245
go fake version 42`

				So(result, ShouldEqual, expected)
			})
		})
	})
}
