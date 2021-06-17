package pacificatlanticwaterflow

import (
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestPaciticAltantic(t *testing.T) {
	type args struct {
		heights string
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
				heights: "[[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]",
			},
			want: want{
				ret: "[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]",
			},
		},
		{
			name: "Example 2",
			args: args{
				heights: "[[2,1],[1,2]]",
			},
			want: want{
				ret: "[[0,0],[0,1],[1,0],[1,1]]",
			},
		},
	}
	for _, tt := range tests {
		got := pacificAtlantic(Matrix(tt.args.heights))
		if !MatrixEqual(SortMatrix(got), SortMatrix(Matrix(tt.want.ret))) {
			t.Errorf("TestPaciticAltantic %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
