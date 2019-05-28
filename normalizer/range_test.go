package normalizer

import (
	"fmt"
	"testing"

	"github.com/guiyomh/charlatan/contract"
	"github.com/stretchr/testify/assert"
)

func TestRange_CanDenormalize(t *testing.T) {
	type args struct {
		ref string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testing a list",
			args: args{"user_{bob,alice}"},
			want: false,
		},
		{
			name: "testing a range reference",
			args: args{"user_{1..3}"},
			want: true,
		},
		{
			name: "testing a range reference with step",
			args: args{"user_{1..3,2}"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{}
			got := r.CanDenormalize(tt.args.ref)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRange_BuildsId(t *testing.T) {
	type args struct {
		match map[string]string
	}
	type want struct {
		err    error
		result []string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "testing a simple range",
			args: args{map[string]string{
				"refbase": "user_",
				"from":    "1",
				"to":      "4",
			}},
			want: want{nil, []string{
				"user_1",
				"user_2",
				"user_3",
				"user_4",
			}},
		},
		{
			name: "testing a range with step",
			args: args{map[string]string{
				"refbase": "user_",
				"from":    "1",
				"to":      "6",
				"step":    "2",
			}},
			want: want{nil, []string{
				"user_1",
				"user_3",
				"user_5",
			}},
		},
		{
			name: "testing a map with missing refbase index",
			args: args{map[string]string{
				"from": "1",
				"to":   "6",
			}},
			want: want{fmt.Errorf("Could not retrieve 'refbase' in 'map[from:1 to:6]'"), make([]string, 0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Range{}
			got, err := r.BuildIds(tt.args.match)
			if tt.want.err != nil {
				assert.EqualError(t, err, fmt.Sprintf("%s", tt.want.err))
			} else {
				assert.NoError(t, err)
			}
			assert.EqualValues(t, tt.want.result, got)
		})
	}
}

func TestRange_Denormalize(t *testing.T) {
	r := &Range{}
	ref := "user_{1..3}"
	fields := map[string]interface{}{
		"username": "<Username()>",
		"password": "<Password()>",
	}
	var bag contract.Bager = make(FixtureBag, 0)
	bag, err := r.Denormalize(bag, ref, fields)
	assert.NoError(t, err)
	assert.Len(t, bag.(FixtureBag), 3)
	assert.True(t, bag.Has("user_1"))
	assert.True(t, bag.Has("user_2"))
	assert.True(t, bag.Has("user_3"))
	_, err = bag.Get("user_1")
	assert.NoError(t, err)
	_, err = bag.Get("user_2")
	assert.NoError(t, err)
	assert.NoError(t, err)
	_, err = bag.Get("user_3")
	assert.NoError(t, err)
}
