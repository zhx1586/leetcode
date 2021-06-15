package partitionequalsubsetsum

import (
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestCanPartition(t *testing.T) {
	type args struct {
		nums string
	}
	type want struct {
		ret bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				nums: "[1,5,11,5]",
			},
			want: want{
				ret: true,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: "[1,2,3,5]",
			},
			want: want{
				ret: false,
			},
		},
		{
			name: "Failed 1",
			args: args{
				nums: "[100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,100,99,97]",
			},
			want: want{
				ret: false,
			},
		},
	}
	for _, tt := range tests {
		got := canPartition(IntSlice(tt.args.nums))
		if got != tt.want.ret {
			t.Errorf("TestCanPartition %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
