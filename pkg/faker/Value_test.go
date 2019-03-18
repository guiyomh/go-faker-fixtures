package faker

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_argConverter(t *testing.T) {
	type args struct {
		value  string
		typeof string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Testing a integer", args{"13", "int"}, 13},
		{"Testing a bool True", args{"True", "bool"}, true},
		{"Testing a bool true", args{"true", "bool"}, true},
		{"Testing a bool 1", args{"1", "bool"}, true},
		{"Testing a bool False", args{"False", "bool"}, false},
		{"Testing a bool false", args{"false", "bool"}, false},
		{"Testing a bool 0", args{"0", "bool"}, false},
		{"Testing a float", args{"3.14", "float"}, "3.14"},
	}
	for _, tt := range tests {
		got := argConverter(tt.args.value, tt.args.typeof)
		assert.Equal(t, tt.want, got)
	}
}

func TestValue_convert(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	type testStruct struct {
		val1 int
		val2 string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test 1", args{val: reflect.ValueOf(true)}, true},
		{"test 2", args{val: reflect.ValueOf(13)}, "13"},
		{"test 3", args{val: reflect.ValueOf("foo")}, "foo"},
		{"test 4", args{val: reflect.ValueOf(testStruct{val1: 16, val2: "foo"})}, testStruct{val1: 16, val2: "foo"}},
		{"test 5", args{val: reflect.ValueOf(func() string { return "foo" })}, "nothing"},
	}
	g := &value{}
	for _, tt := range tests {
		got := g.convert(tt.args.val)
		assert.Equal(t, tt.want, got)
	}
}
