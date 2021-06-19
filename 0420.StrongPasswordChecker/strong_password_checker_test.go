package strongpasswordchecker

import (
	"reflect"
	"testing"
)

func TestStrongPasswordChecker(t *testing.T) {
	type args struct {
		password string
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
				password: "a",
			},
			want: want{
				ret: 5,
			},
		},
		{
			name: "Example 2",
			args: args{
				password: "aA1",
			},
			want: want{
				ret: 3,
			},
		},
		{
			name: "Example 3",
			args: args{
				password: "1337C0d3",
			},
			want: want{
				ret: 0,
			},
		},
		{
			name: "Failed 1",
			args: args{
				password: "aaa111",
			},
			want: want{
				ret: 2,
			},
		},
	}
	for _, tt := range tests {
		got := strongPasswordChecker(tt.args.password)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestStrongPasswordChecker %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
