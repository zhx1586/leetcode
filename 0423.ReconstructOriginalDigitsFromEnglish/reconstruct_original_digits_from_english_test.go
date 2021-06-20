package reconstructoriginaldigitsfromenglish

import (
	"reflect"
	"testing"
)

func TestOriginalDigits(t *testing.T) {
	type args struct {
		s string
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
				s: "owoztneoer",
			},
			want: want{
				ret: "012",
			},
		},
		{
			name: "Example 2",
			args: args{
				s: "fviefuro",
			},
			want: want{
				ret: "45",
			},
		},
		{
			name: "Failed 1",
			args: args{
				s: "esnve",
			},
			want: want{
				ret: "7",
			},
		},
	}
	for _, tt := range tests {
		got := originalDigits(tt.args.s)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestOriginalDigits %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
