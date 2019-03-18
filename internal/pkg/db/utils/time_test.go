package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_tryStrToDate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{"test 1", args{s: "19/03/2019"}, time.Date(2019, time.March, 19, 0, 0, 0, 0, time.Local), false},
		{"test 2", args{s: "2019-03-19"}, time.Date(2019, time.March, 19, 0, 0, 0, 0, time.Local), false},
		{"test 3", args{s: "2019-03-19 15:23:55"}, time.Date(2019, time.March, 19, 15, 23, 55, 0, time.Local), false},
		{"test 4", args{s: "2019-03-19 13:55"}, time.Date(2019, time.March, 19, 13, 55, 0, 0, time.Local), false},
		{"test 5", args{s: "20190319"}, time.Date(2019, time.March, 19, 0, 0, 0, 0, time.Local), false},
		{"test 6", args{s: "20190319 12:06"}, time.Date(2019, time.March, 19, 12, 06, 0, 0, time.Local), false},
		{"test 6", args{s: "foo bar"}, time.Time{}, true},
	}
	for _, tt := range tests {
		got, err := TryStrToDate(tt.args.s)
		assert.Equal(t, tt.wantErr, err != nil)
		assert.Equal(t, tt.want, got)
	}
}
