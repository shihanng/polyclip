package polygon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClip(t *testing.T) {
	type args struct {
		polyA Polygon
		polyB Polygon
	}
	tests := []struct {
		name string
		args args
		want Polygon
	}{
		{
			name: "normal",
			args: args{
				polyA: Polygon{
					Vertices: []Point{
						{2, 1},
						{5, 1},
						{5, 4},
						{2, 4},
					}},
				polyB: Polygon{
					Vertices: []Point{
						{4, 3},
						{7, 3},
						{7, 6},
						{4, 6},
					},
				},
			},
			want: Polygon{
				Vertices: []Point{
					{4, 4},
					{4, 3},
					{5, 3},
					{5, 4},
				},
			},
		},
		{
			name: "same polygons",
			args: args{
				polyA: Polygon{
					Vertices: []Point{
						{2, 1},
						{5, 1},
						{5, 4},
						{2, 4},
					},
				},
				polyB: Polygon{
					Vertices: []Point{
						{2, 1},
						{5, 1},
						{5, 4},
						{2, 4},
					},
				},
			},
			want: Polygon{
				Vertices: []Point{
					{2, 4},
					{2, 1},
					{5, 1},
					{5, 4},
				},
			},
		},
		{
			name: "irregular polygons",
			args: args{
				polyA: Polygon{
					Vertices: []Point{
						{-2.76, 0.91},
						{-3.06, -1.81},
						{-1.24, -2.37},
						{0.74, -2.13},
						{1.6, 0.51},
						{0.88, 1.73},
						{-0.96, 2.23},
					},
				},
				polyB: Polygon{
					Vertices: []Point{
						{-2.24, 1.65},
						{-3.14, -0.53},
						{-2.22, -2.41},
						{0.24, -2.89},
						{1.98, -1.31},
						{0.28, 2.47},
					},
				},
			},
			want: Polygon{
				Vertices: []Point{
					{-0.6862062987575844, 2.1555995377058657},
					{-1.3607782101167338, 1.9360959792477295},
					{-2.4523684210526318, 1.1355964912280703},
					{-2.83819397993311, 0.2010412486064674},
					{-2.9595043047221496, -0.8988390294808247},
					{-2.416763005780347, -2.0079190751445077},
					{-1.2399999999999984, -2.37},
					{0.7400000000000001, -2.1299999999999994},
					{1.4157936631002224, -0.055470615599317916},
					{0.5756037476249757, 1.8127163729279956},
				},
			},
		},
		{
			name: "different orientation",
			args: args{
				polyA: Polygon{
					Vertices: []Point{
						{2, 1},
						{5, 1},
						{5, 4},
						{2, 4},
					},
				},
				polyB: Polygon{
					Vertices: []Point{
						{4, 6},
						{7, 6},
						{7, 3},
						{4, 3},
					},
				},
			},
			want: Polygon{
				Vertices: []Point{
					{4, 4},
					{5, 4},
					{5, 3},
					{4, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, &tt.want, Clip(&tt.args.polyA, &tt.args.polyB))
		})
	}
}
