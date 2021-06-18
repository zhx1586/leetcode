package battleshipsinaboard

import (
	"reflect"
	"testing"
)

func TestCountBattleships(t *testing.T) {
	type args struct {
		board [][]byte
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
				board: [][]byte{
					{'X', '.', '.', 'X'},
					{'.', '.', '.', 'X'},
					{'.', '.', '.', 'X'},
				},
			},
			want: want{
				ret: 2,
			},
		},
		{
			name: "Example 2",
			args: args{
				board: [][]byte{{'.'}},
			},
			want: want{
				ret: 0,
			},
		},
	}
	for _, tt := range tests {
		got := countBattleships(tt.args.board)
		if !reflect.DeepEqual(got, tt.want.ret) {
			t.Errorf("TestCountBattleships %s failed: (got: %+v, want: %+v)",
				tt.name, got, tt.want.ret)
		}
	}
}
