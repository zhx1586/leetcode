package maximumxoroftwonumbersinanarray

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestFindMaximumXOR(t *testing.T) {
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
				nums: "[3,10,5,25,2,8]",
			},
			want: want{
				ret: 28,
			},
		},
		{
			name: "Example 2",
			args: args{
				nums: "[0]",
			},
			want: want{
				ret: 0,
			},
		},
		{
			name: "Example 3",
			args: args{
				nums: "[2,4]",
			},
			want: want{
				ret: 6,
			},
		},
		{
			name: "Example 4",
			args: args{
				nums: "[8,10,2]",
			},
			want: want{
				ret: 10,
			},
		},
		{
			name: "Example 5",
			args: args{
				nums: "[14,70,53,83,49,91,36,80,92,51,66,70]",
			},
			want: want{
				ret: 127,
			},
		},
		{
			name: "Failed 1",
			args: args{
				nums: "[1]",
			},
			want: want{
				ret: 0,
			},
		},
		{
			name: "Failed 2",
			args: args{
				nums: "[4,6,7]",
			},
			want: want{
				ret: 3,
			},
		},
	}
	for _, tt := range tests {
		got := findMaximumXOR(IntSlice(tt.args.nums))
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestFindMaximumXOR %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
