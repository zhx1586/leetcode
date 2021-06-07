package frogjump

import (
	"reflect"
	"testing"
)

func TestCanCross(t *testing.T) {
	type args struct {
		stones []int
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
				stones: []int{0, 1, 3, 5, 6, 8, 12, 17},
			},
			want: want{
				ret: true,
			},
		},
		{
			name: "Example 2",
			args: args{
				stones: []int{0, 1, 2, 3, 4, 8, 9, 11},
			},
			want: want{
				ret: false,
			},
		},
	}
	for _, tt := range tests {
		got := canCross(tt.args.stones)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestCanCross %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
