package arithmeticslices

import (
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
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
				nums: "[1,2,3,4]",
			},
			want: want{
				ret: 3,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: "[1]",
			},
			want: want{
				ret: 0,
			},
		},
		{
			name: "Customized 1",
			args: args{
				nums: "[1,2,3,4,1,2,3,4]",
			},
			want: want{
				ret: 6,
			},
		},
	}
	for _, tt := range tests {
		got := numberOfArithmeticSlices(IntSlice(tt.args.nums))
		if got != tt.want.ret {
			t.Errorf("TestNumberOfArithmeticSlices %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
