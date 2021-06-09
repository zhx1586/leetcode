package convertanumbertohexadecimal

import (
	"reflect"
	"testing"
)

func TestToHex(t *testing.T) {
	type args struct {
		num int
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
				num: 26,
			},
			want: want{
				ret: "1a",
			},
		},
		{
			name: "Example 2",
			args: args{
				num: -1,
			},
			want: want{
				ret: "ffffffff",
			},
		},
		{
			name: "Failed 1",
			args: args{
				num: 0,
			},
			want: want{
				ret: "0",
			},
		},
	}
	for _, tt := range tests {
		got := toHex(tt.args.num)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestToHex %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
