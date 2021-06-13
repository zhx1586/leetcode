package splitarraylargestsum

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestSplitArray(t *testing.T) {
	type args struct {
		nums string
		m    int
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
				nums: "[7,2,5,10,8]",
				m:    2,
			},
			want: want{
				ret: 18,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: "[1,2,3,4,5]",
				m:    2,
			},
			want: want{
				ret: 9,
			},
		},
		{
			name: "Example 3",
			args: args{
				nums: "[1,4,4]",
				m:    3,
			},
			want: want{
				ret: 4,
			},
		},
	}
	for _, tt := range tests {
		got := splitArray(IntSlice(tt.args.nums), tt.args.m)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestSplitArray %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
