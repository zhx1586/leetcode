package addstrings

import (
	"testing"
)

func TestAddStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
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
				num1: "11",
				num2: "123",
			},
			want: want{
				ret: "134",
			},
		},
		{
			name: "Example 2",
			args: args{
				num1: "456",
				num2: "77",
			},
			want: want{
				ret: "533",
			},
		},
		{
			name: "Example 3",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: want{
				ret: "0",
			},
		},
		{
			name: "Customized 1",
			args: args{
				num1: "9999",
				num2: "1",
			},
			want: want{
				ret: "10000",
			},
		},
	}
	for _, tt := range tests {
		got := addStrings(tt.args.num1, tt.args.num2)
		if got != tt.want.ret {
			t.Errorf("TestAddStrings %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
