package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRelation(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *relation
		wantErr bool
	}{
		{"test 1", args{value: "@bar"}, &relation{recordName: "bar"}, false},
		{"test 2", args{value: "@foo.bar"}, &relation{recordName: "foo", fieldName: "bar"}, false},
		{"test 3", args{value: "@foo_bar123.field_3"}, &relation{recordName: "foo_bar123", fieldName: "field_3"}, false},
		//{"test 4", args{value: "bar.foo"}, nil, true},
	}
	for _, tt := range tests {
		got, err := NewRelation(tt.args.value)
		if true == tt.wantErr && assert.Error(t, err) {
			assert.Equal(t, errNotFoundRelation, err)
		}
		assert.Equal(t, tt.want, got)
	}
}
