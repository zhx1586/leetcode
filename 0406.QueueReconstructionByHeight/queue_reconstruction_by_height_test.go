package queuereconstructionbyheight

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestReconstructQueue(t *testing.T) {
	type args struct {
		people string
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
				people: "[[7,0],[4,4],[7,1],[5,0],[6,1],[5,2]]",
			},
			want: want{
				ret: "[[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]",
			},
		},
		{
			name: "Example 2",
			args: args{
				people: "[[6,0],[5,0],[4,0],[3,2],[2,2],[1,4]]",
			},
			want: want{
				ret: "[[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]",
			},
		},
	}
	for _, tt := range tests {
		got := reconstructQueue(Matrix(tt.args.people))
		if !reflect.DeepEqual(got, Matrix(tt.want.ret)) {
			t.Errorf("TestReconstructQueue %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
