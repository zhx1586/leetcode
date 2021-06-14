package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	type want struct {
		ret []string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Example 1",
			args: args{
				n: 3,
			},
			want: want{
				ret: []string{
					"1", "2", "Fizz",
				},
			},
		},
		{
			name: "Example 2",
			args: args{
				n: 5,
			},
			want: want{
				ret: []string{
					"1", "2", "Fizz", "4", "Buzz",
				},
			},
		},
		{
			name: "Example 3",
			args: args{
				n: 15,
			},
			want: want{
				ret: []string{
					"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz",
				},
			},
		},
	}
	for _, tt := range tests {
		got := fizzBuzz(tt.args.n)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestFizzBuzz %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
