package sumofleftleaves

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestSumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
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
				root: NewBinaryTree("[3,9,20,null,null,15,7]"),
			},
			want: want{
				ret: 24,
			},
		},
		{
			name: "Example 2",
			args: args{
				root: NewBinaryTree("[1]"),
			},
			want: want{
				ret: 0,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		got := sumOfLeftLeaves(tt.args.root)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestSumOfLeftLeaves %s failed: (got:%+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
