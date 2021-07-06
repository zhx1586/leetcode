package findrightinterval

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestFindRightInterval(t *testing.T) {
	type args struct {
		intervals string
	}
	type want struct {
		ret string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				intervals: "[[1,2]]",
			},
			want: want{
				ret: "[-1]",
			},
		},
		{
			name: "Example 2",
			args: args{
				intervals: "[[3,4],[2,3],[1,2]]",
			},
			want: want{
				ret: "[-1,0,1]",
			},
		},
		{
			name: "Example 3",
			args: args{
				intervals: "[[1,4],[2,3],[3,4]]",
			},
			want: want{
				ret: "[-1,2,-1]",
			},
		},
	}
	for _, tt := range tests {
		got := findRightInterval(Matrix(tt.args.intervals))
		if !reflect.DeepEqual(got, IntSlice(tt.want.ret)) {
			t.Errorf("TestFindRightInterval %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
