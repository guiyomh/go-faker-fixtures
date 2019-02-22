package loader

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/afero"
)

func initLoaderFs() afero.Fs {
	fs := afero.NewMemMapFs()

	fs.MkdirAll("fixtures/", 0755)
	afero.WriteFile(fs, "fixtures/tableA.yml", []byte("table A"), 0644)
	afero.WriteFile(fs, "fixtures/tableB.yaml", []byte("table B"), 0644)
	return fs
}

func TestLoad(t *testing.T) {
	Convey("Given I have a filesystem with fixtures files", t, func() {
		fs := initLoaderFs()
		l := NewFileLoader(fs)
		fixtures := []string{
			"fixtures/tableA.yml",
			"fixtures/tableB.yaml",
		}
		Convey("When I load the fixtures", func() {
			got := l.Load(fixtures)
			Convey("Then I should have the fixtures's content", func() {
				expected := []byte("table A\ntable B")
				So(got, ShouldResemble, expected)
			})
		})
	})
}

func TestParseError(t *testing.T) {
	Convey("Given I have a filesystem with fixtures files", t, func() {
		fs := initLoaderFs()
		l := NewFileLoader(fs)
		fixtures := []string{
			"fixtures/tableC.yml",
		}
		Convey("When I load the fixtures", func() {
			got := l.Load(fixtures)
			Convey("Then I should have the fixtures's content", func() {
				var expected []byte
				So(got, ShouldEqual, expected)
			})
		})
	})
}
