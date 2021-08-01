package pathsumiii

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestPathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
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
				root:      NewBinaryTree("[10,5,-3,3,2,null,11,3,-2,null,1]"),
				targetSum: 8,
			},
			want: want{
				ret: 3,
			},
		},
		{
			name: "Example 2",
			args: args{
				root:      NewBinaryTree("[5,4,8,11,null,13,4,7,2,null,null,5,1]"),
				targetSum: 22,
			},
			want: want{
				ret: 3,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		got := pathSum(tt.args.root, tt.args.targetSum)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestPathSum %s failed: (got:%+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
