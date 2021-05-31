package binarywatch

import (
	"reflect"
	"sort"
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	type args struct {
		turnedOn int
	}
	type want struct {
		ret []string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				turnedOn: 1,
			},
			want: want{
				ret: []string{
					"0:01", "0:02", "0:04", "0:08", "0:16", "0:32",
					"1:00", "2:00", "4:00", "8:00",
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				turnedOn: 9,
			},
			want: want{
				ret: []string{},
			},
		},
	}
	for _, tt := range tests {
		got := readBinaryWatch(tt.args.turnedOn)
		sort.Strings(got)
		sort.Strings(tt.want.ret)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestReadBinaryWatch %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
