package longestpalindrome

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		s string
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
				s: "abccccdd",
			},
			want: want{
				ret: 7,
			},
		},
		{
			name: "Example 2",
			args: args{
				s: "a",
			},
			want: want{
				ret: 1,
			},
		},
		{
			name: "Example 3",
			args: args{
				s: "bb",
			},
			want: want{
				ret: 2,
			},
		},
	}
	for _, tt := range tests {
		got := longestPalindrome(tt.args.s)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestLongestPalindrome %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
