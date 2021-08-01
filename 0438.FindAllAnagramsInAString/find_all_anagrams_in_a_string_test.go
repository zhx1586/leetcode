package findallanagramsinastring

import (
	"reflect"
	"testing"

	. "github.com/zhx1586/leetcode/utils"
)

func TestFindAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	type want struct {
		ret []int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			want: want{
				ret: IntSlice("[0,6]"),
			},
		},
		{
			name: "Example 2",
			args: args{
				s: "abab",
				p: "ab",
			},
			want: want{
				ret: IntSlice("[0,1,2]"),
			},
		},
	}
	for _, tt := range tests {
		got := findAnagrams(tt.args.s, tt.args.p)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestFindAnagrams %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
