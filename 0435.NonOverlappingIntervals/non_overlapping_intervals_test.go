package nonoverlappingintervals

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestEraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals string
	}
	type want struct {
		ret int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				intervals: "[[1,2],[2,3],[3,4],[1,3]]",
			},
			want: want{
				ret: 1,
			},
		},
		{
			name: "Example 2",
			args: args{
				intervals: "[[1,2],[1,2],[1,2]]",
			},
			want: want{
				ret: 2,
			},
		},
		{
			name: "Example 3",
			args: args{
				intervals: "[[1,2],[2,3]]",
			},
			want: want{
				ret: 0,
			},
		},
	}
	for _, tt := range tests {
		got := eraseOverlapIntervals(Matrix(tt.args.intervals))
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestEraseOverlapIntervals %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
