package longestrepeatingcharacterreplacement

import (
	"reflect"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
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
				s: "ABAB",
				k: 2,
			},
			want: want{
				ret: 4,
			},
		},
		{
			name: "Example 2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			want: want{
				ret: 4,
			},
		},
		{
			name: "Failed 1",
			args: args{
				s: "ABAA",
				k: 0,
			},
			want: want{
				ret: 2,
			},
		},
	}
	for _, tt := range tests {
		got := characterReplacement(tt.args.s, tt.args.k)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestCharacterReplacement %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
