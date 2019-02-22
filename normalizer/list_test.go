package normalizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_CanDenormalize(t *testing.T) {
	type args struct {
		ref string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testing a good list",
			args: args{"user_{bob,alice}"},
			want: true,
		},
		{
			name: "testing a range reference",
			args: args{"user_{1..3}"},
			want: false,
		},
		{
			name: "testing a range reference with step",
			args: args{"user_{1..3,2}"},
			want: false,
		},
		{
			name: "testing a list with space and digit",
			args: args{"user_{bob, alice,maurice7}"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{}
			got := l.CanDenormalize(tt.args.ref)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_BuildsId(t *testing.T) {
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
			name: "testing a good map",
			args: args{map[string]string{
				"list":    "bob,alice",
				"refbase": "user_",
			}},
			want: want{nil, []string{
				"user_bob",
				"user_alice",
			}},
		},
		{
			name: "testing a map with missing list index",
			args: args{map[string]string{
				"refbase": "user_",
			}},
			want: want{fmt.Errorf("Could not retrieve 'list' in 'map[refbase:user_]'"), make([]string, 0)},
		},
		{
			name: "testing a map with missing refbase index",
			args: args{map[string]string{
				"list": "bob,alice",
			}},
			want: want{fmt.Errorf("Could not retrieve 'refbase' in 'map[list:bob,alice]'"), make([]string, 0)},
		},
		{
			name: "testing a good map with space",
			args: args{map[string]string{
				"list":    "bob,  alice",
				"refbase": "user_",
			}},
			want: want{nil, []string{
				"user_bob",
				"user_alice",
			}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			l := &List{}
			got, err := l.BuildIds(tt.args.match)

			if tt.want.err != nil {
				assert.EqualError(t, err, fmt.Sprintf("%s", tt.want.err))
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want.result, got)
		})
	}
}

func TestList_Denormalize(t *testing.T) {
	l := &List{}
	ref := "user_{bob,alice}"
	fields := map[string]interface{}{
		"username": "<Username()>",
		"password": "<Password()>",
	}
	bag, err := l.Denormalize(ref, fields)
	assert.NoError(t, err)
	assert.Len(t, bag.(FixtureBag), 2)
	assert.True(t, bag.Has("user_bob"))
	assert.True(t, bag.Has("user_alice"))
	_, err = bag.Get("user_bob")
	assert.NoError(t, err)
	_, err = bag.Get("user_alice")
	assert.NoError(t, err)
}
