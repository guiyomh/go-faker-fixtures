package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveToJSON(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		args  args
		wantR interface{}
	}{
		{
			args{
				v: map[interface{}]interface{}{
					"2":   true,
					"foo": "bar",
				},
			},
			jsonMap{
				"2":   true,
				"foo": "bar",
			},
		},
		{
			args{
				v: []interface{}{
					true,
					"bar",
					10,
				},
			},
			jsonArray{
				true,
				"bar",
				10,
			},
		},
		{
			args{
				v: map[interface{}]interface{}{
					"2": true,
					"foo": map[interface{}]interface{}{
						"biloute": "stuff",
						"dragons": []interface{}{
							"Abraxas",
							"Movoss, Destroyer Of Men",
							"Xordasdig, Lord Of The Black",
						},
					},
				},
			},
			jsonMap{
				"2": true,
				"foo": jsonMap{
					"biloute": "stuff",
					"dragons": jsonArray{
						"Abraxas",
						"Movoss, Destroyer Of Men",
						"Xordasdig, Lord Of The Black",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.wantR, RecursiveToJSON(tt.args.v))
	}
}
