package polyclip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClip(t *testing.T) {
	type args struct {
		polyA []Point
		polyB []Point
	}
	tests := []struct {
		name string
		args args
		want []Point
	}{
		{
			name: "normal",
			args: args{
				polyA: []Point{
					{2, 1},
					{5, 1},
					{5, 4},
					{2, 4},
				},
				polyB: []Point{
					{4, 3},
					{7, 3},
					{7, 6},
					{4, 6},
				},
			},
			want: []Point{
				{4, 4},
				{4, 3},
				{5, 3},
				{5, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Clip(tt.args.polyA, tt.args.polyB))
		})
	}
}
