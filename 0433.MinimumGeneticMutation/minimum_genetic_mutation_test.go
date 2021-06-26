package minimumgeneticmutation

import (
	"reflect"
	"testing"
)

func TestMinMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
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
				start: "AACCGGTT",
				end:   "AACCGGTA",
				bank: []string{
					"AACCGGTA",
				},
			},
			want: want{
				ret: 1,
			},
		},
		{
			name: "Example 2",
			args: args{
				start: "AACCGGTT",
				end:   "AAACGGTA",
				bank: []string{
					"AACCGGTA",
					"AACCGCTA",
					"AAACGGTA",
				},
			},
			want: want{
				ret: 2,
			},
		},
		{
			name: "Example 3",
			args: args{
				start: "AAAAACCC",
				end:   "AACCCCCC",
				bank: []string{
					"AAAACCCC",
					"AAACCCCC",
					"AACCCCCC",
				},
			},
			want: want{
				ret: 3,
			},
		},
	}
	for _, tt := range tests {
		got := minMutation(tt.args.start, tt.args.end, tt.args.bank)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestMinMutation %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
