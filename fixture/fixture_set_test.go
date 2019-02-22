package fixture

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_newFixtureSet(t *testing.T) {
	type args struct {
		name string
		data map[string]interface{}
	}

	want := fixtureSet{
		Name: "bob_user",
		Fields: map[string]field{
			"username": field{"username", "bob"},
			"age":      field{"age", 20},
			"isActive": field{"isActive", true},
			"birthDay": field{"birthDay", func() time.Time {
				at, _ := time.Parse("01/12/2016", "29/04/1982")
				return at
			}()},
		},
	}

	Convey("Given i have loaded data from yaml", t, func() {
		a := args{
			name: "bob_user",
			data: map[string]interface{}{
				"username": "bob",
				"age":      20,
				"isActive": true,
				"birthDay": func() time.Time {
					at, _ := time.Parse("01/12/2016", "29/04/1982")
					return at
				}(),
			},
		}
		Convey("When i a create a fixture set", func() {
			got := newFixtureSet(a.name, a.data)
			So(got, ShouldHaveSameTypeAs, fixtureSet{})
			Convey("Then the set must containe the same value of the fixture set", func() {
				So(got, ShouldResemble, want)
				So(got.Name, ShouldEqual, "bob_user")
				So(len(got.Fields), ShouldEqual, 4)
			})
		})
	})

}

func Test_newFixtureSets(t *testing.T) {
	type args struct {
		name string
		data map[string]map[string]interface{}
	}

	want := map[string]fixtureSet{
		"bob": fixtureSet{
			Name: "bob",
			Fields: map[string]field{
				"username": field{"username", "bob"},
				"age":      field{"age", 20},
				"isActive": field{"isActive", true},
				"birthDay": field{"birthDay", func() time.Time {
					at, _ := time.Parse("01/12/2016", "29/04/1982")
					return at
				}()},
			},
		},
		"bobette": fixtureSet{
			Name: "bobette",
			Fields: map[string]field{
				"username": field{"username", "bobette"},
				"age":      field{"age", 10},
				"isActive": field{"isActive", false},
				"birthDay": field{"birthDay", func() time.Time {
					at, _ := time.Parse("01/12/2016", "29/04/2009")
					return at
				}()},
			},
		},
	}

	Convey("Given i have loaded data from yaml", t, func() {
		a := args{
			name: "user",
			data: map[string]map[string]interface{}{
				"bob": {
					"username": "bob",
					"age":      20,
					"isActive": true,
					"birthDay": func() time.Time {
						at, _ := time.Parse("01/12/2016", "29/04/1982")
						return at
					}(),
				},
				"bobette": {
					"username": "bobette",
					"age":      10,
					"isActive": false,
					"birthDay": func() time.Time {
						at, _ := time.Parse("01/12/2016", "29/04/2009")
						return at
					}(),
				},
			},
		}
		Convey("When i a create a fixture set", func() {
			got := newFixtureSets(a.data)
			So(got, ShouldHaveSameTypeAs, map[string]fixtureSet{})
			Convey("Then the set must containe the same value of the fixture set", func() {
				So(got, ShouldResemble, want)
			})
			Convey("And the length of the set must be equal to 2", func() {
				So(len(got), ShouldEqual, 2)
			})
			Convey("And the keys of the fixture set must equal to [bob, bobette]", func() {
				So(got, ShouldContainKey, "bob")
				So(got, ShouldContainKey, "bobette")
			})
		})
	})
}
