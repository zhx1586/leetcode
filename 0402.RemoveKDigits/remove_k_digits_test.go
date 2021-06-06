package removekdigits

import (
	"reflect"
	"testing"
)

func TestRemoveKDigits(t *testing.T) {
	type args struct {
		num string
		k   int
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
				num: "1432219",
				k:   3,
			},
			want: want{
				ret: "1219",
			},
		},
		{
			name: "Example 2",
			args: args{
				num: "10200",
				k:   1,
			},
			want: want{
				ret: "200",
			},
		},
		{
			name: "Example 3",
			args: args{
				num: "10",
				k:   2,
			},
			want: want{
				ret: "0",
			},
		},
		{
			name: "Failed 1",
			args: args{
				num: "112",
				k:   1,
			},
			want: want{
				ret: "11",
			},
		},
	}
	for _, tt := range tests {
		got := removeKdigits(tt.args.num, tt.args.k)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestRemoveKDigits %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
