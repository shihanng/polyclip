package polygon

// Point describes a vertex of a polygon in Cartesian coordinates.
type Point struct {
	X, Y float64
}

// Clip computes the clipping polygon based on the Sutherland–Hodgman algorithm.
func Clip(polyA, polyB []Point) []Point {
	output := polyB
	a1 := polyA[len(polyA)-1]
	for _, a2 := range polyA { // Clipping edge: [cp1, cp2]
		input := output
		output = nil
		b1 := input[len(input)-1]
		for _, b2 := range input {
			if isInside(a1, a2, b2) {
				if !isInside(a1, a2, b1) {
					output = append(output, intersection(a1, a2, b1, b2))
				}
				output = append(output, b2)
			} else if isInside(a1, a2, b1) {
				output = append(output, intersection(a1, a2, b1, b2))
			}
			b1 = b2
		}
		a1 = a2
	}
	return output
}

func isInside(a1, a2, p Point) bool {
	return (a2.X-a1.X)*(p.Y-a1.Y) > (a2.Y-a1.Y)*(p.X-a1.X)
}

func intersection(a1, a2, b1, b2 Point) Point {
	dcx, dcy := a1.X-a2.X, a1.Y-a2.Y
	dpx, dpy := b1.X-b2.X, b1.Y-b2.Y
	n1 := a1.X*a2.Y - a1.Y*a2.X
	n2 := b1.X*b2.Y - b1.Y*b2.X
	n3 := 1 / (dcx*dpy - dcy*dpx)

	return Point{
		X: (n1*dpx - n2*dcx) * n3,
		Y: (n1*dpy - n2*dcy) * n3,
	}
}
