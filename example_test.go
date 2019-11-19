package polygon

import (
	"fmt"
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func ExampleClip() {
	polyA := []Point{
		{-2.76, 0.91},
		{-3.06, -1.81},
		{-1.24, -2.37},
		{0.74, -2.13},
		{1.6, 0.51},
		{0.88, 1.73},
		{-0.96, 2.23},
	}

	polyB := []Point{
		{-2.24, 1.65},
		{-3.14, -0.53},
		{-2.22, -2.41},
		{0.24, -2.89},
		{1.98, -1.31},
		{0.28, 2.47},
	}

	clipped := Clip(polyA, polyB)

	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	pA := mustToXYer(polyA)
	pA.LineStyle.Color = color.NRGBA{R: 255, G: 0, B: 0, A: 255}

	pB := mustToXYer(polyB)
	pB.LineStyle.Color = color.NRGBA{R: 0, G: 0, B: 255, A: 255}

	pClipped := mustToXYer(clipped)
	pClipped.LineStyle.Width = 0
	pClipped.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}

	p.Add(pA, pB, pClipped)

	if err = p.Save(200, 200, "./testdata/polygons_clipping.png"); err != nil {
		log.Panic(err)
	}

	fmt.Println(pClipped)
	// Output: &{[[{-0.6862062987575844 2.1555995377058657} {-1.3607782101167338 1.9360959792477295} {-2.4523684210526318 1.1355964912280703} {-2.83819397993311 0.2010412486064674} {-2.9595043047221496 -0.8988390294808247} {-2.416763005780347 -2.0079190751445077} {-1.2399999999999984 -2.37} {0.7400000000000001 -2.1299999999999994} {1.4157936631002224 -0.055470615599317916} {0.5756037476249757 1.8127163729279956}]] {{0} 0 [] 0} {0 255 0 255}}
}

func mustToXYer(points []Point) *plotter.Polygon {
	pp := make([]plotter.XY, 0, len(points))

	for _, p := range points {
		pp = append(pp, plotter.XY{
			X: p.X,
			Y: p.Y,
		})
	}

	polygon, err := plotter.NewPolygon(plotter.XYs(pp))
	if err != nil {
		log.Panic(err)
	}

	return polygon
}
