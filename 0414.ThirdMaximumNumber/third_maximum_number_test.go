package thirdmaximumnumber

import (
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestThirdMax(t *testing.T) {
	type args struct {
		nums string
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
				nums: "[3,2,1]",
			},
			want: want{
				ret: 1,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: "[1,2]",
			},
			want: want{
				ret: 2,
			},
		},
		{
			name: "Example 3",
			args: args{
				nums: "[2,2,3,1]",
			},
			want: want{
				ret: 1,
			},
		},
		{
			name: "Customized 1",
			args: args{
				nums: "[2,2,2,3,3,3]",
			},
			want: want{
				ret: 3,
			},
		},
		{
			name: "Failed 1",
			args: args{
				nums: "[1,2,2,5,3,5]",
			},
			want: want{
				ret: 2,
			},
		},
	}
	for _, tt := range tests {
		got := thirdMax(IntSlice(tt.args.nums))
		if got != tt.want.ret {
			t.Errorf("TestThirdMax %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
