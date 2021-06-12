package trappingrainwaterii

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestTrapRainWater(t *testing.T) {
	type args struct {
		heightMap string
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
				heightMap: "[[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]",
			},
			want: want{
				ret: 4,
			},
		},
		{
			name: "Example 2",
			args: args{
				heightMap: "[[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]",
			},
			want: want{
				ret: 10,
			},
		},
		{
			name: "Customized 1",
			args: args{
				heightMap: "[[1,1,1,1,1,1,1,1],[1,1,4,3,1,3,2,1],[1,3,2,1,3,2,4,1],[1,2,3,3,2,3,1,1],[1,1,1,1,1,1,1,1]]",
			},
			want: want{
				ret: 4,
			},
		},
	}
	for _, tt := range tests {
		got := trapRainWater(Matrix(tt.args.heightMap))
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestTrapRainWater %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
