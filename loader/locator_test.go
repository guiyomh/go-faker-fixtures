package loader

import (
	"fmt"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func initFs() afero.Fs {
	fs := afero.NewMemMapFs()

	fs.MkdirAll("fixtures/dev", 0755)
	afero.WriteFile(fs, "fixtures/tableA.yml", []byte("table A"), 0644)
	afero.WriteFile(fs, "fixtures/tableB.yaml", []byte("table B"), 0644)
	afero.WriteFile(fs, "fixtures/tableC.yaml.disable", []byte("table C"), 0644)
	afero.WriteFile(fs, "fixtures/dev/tableD.yaml", []byte("table D"), 0644)
	afero.WriteFile(fs, "fixtures/dev/tableE.yaml", []byte(""), 0644)
	return fs
}

func TestLocateFiles(t *testing.T) {
	Convey("Given I have a file system with fixtures files in \"fixtures/\"", t, func() {
		fs := initFs()
		l := NewFixtureLocator(fs)
		Convey("When I search fixtures files", func() {
			got, err := l.LocateFiles("fixtures")
			So(err, ShouldBeNil)
			goodFile := []string{
				"fixtures/tableA.yml",
				"fixtures/tableB.yaml",
				"fixtures/dev/tableD.yaml",
			}
			for i, f := range goodFile {
				keyword := "And"
				if i == 0 {
					keyword = "Then"
				}
				Convey(fmt.Sprintf("%s I should see a file \"%s\"", keyword, f), func() {
					So(got, ShouldContain, f)
				})
			}
			badFile := []string{
				"fixtures/tableC.yaml.disable",
				"fixtures/dev/tableE.yaml",
			}
			for i, f := range badFile {
				keyword := "And"
				if i == 0 {
					keyword = "Then"
				}
				Convey(fmt.Sprintf("%s I should not see a file \"%s\"", keyword, f), func() {
					So(got, ShouldNotContain, f)
				})
			}
		})
	})
}

func TestIsValidFixture(t *testing.T) {
	fs := initFs()

	f1, _ := fs.Stat("fixtures/tableA.yml")
	f2, _ := fs.Stat("fixtures/tableB.yaml")
	f3, _ := fs.Stat("fixtures/dev/tableD.yaml")
	f4, _ := fs.Stat("fixtures/tableC.yaml.disable")
	f5, _ := fs.Stat("fixtures/dev/tableE.yaml")

	type args struct {
		f os.FileInfo
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test table A",
			args: args{f: f1},
			want: true,
		},
		{
			name: "Test table B",
			args: args{f: f2},
			want: true,
		},
		{
			name: "Test table D",
			args: args{f: f3},
			want: true,
		},
		{
			name: "Test table C",
			args: args{f: f4},
			want: false,
		},
		{
			name: "Test table E",
			args: args{f: f5},
			want: false,
		},
	}
	for _, tt := range tests {
		fmt.Printf("- File %v \n", tt.args.f)
		t.Run(tt.name, func(t *testing.T) {
			got := isValidFixture(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
