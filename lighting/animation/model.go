package animation

import "math"

type Curve []Handle

func (c Curve) Sample(head float64) float64 {
	head = math.Min(head, 1.0)
	head = math.Max(head, 0.0)
	index := int(math.Floor(head * float64(len(c))))

	x := head*float64(len(c)) - float64(index)

	handle := c[index]
	return handle.Sample(x)
}

type Handle struct {
	P0 Point `json:"p0"`
	P1 Point `json:"p1"`
	P2 Point `json:"p2"`
	P3 Point `json:"p3"`
}

func (h Handle) Sample(x float64) float64 {
	x = math.Min(x, 1.0)
	x = math.Max(x, 0.0)
	return NewRasterizedBezierSegment(h, 100).Sample(x)
}

type Animation struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Red        Curve  `json:"red"`
	Green      Curve  `json:"green"`
	Blue       Curve  `json:"blue"`
	Brightness Curve  `json:"brightness"`
}

// Point represents a 2D point.
type Point struct {
	X, Y float64
}

// RasterizedBezier provides a simple, fast, and approximate way to solve for Y given X.
type RasterizedBezier struct {
	points []Point // Our Lookup Table
}

// NewRasterizedBezierSegment creates a solver that uses a lookup table for approximation.
// quality determines how many points to pre-compute. 100 is a good default.
func NewRasterizedBezierSegment(h Handle, resolution int) *RasterizedBezier {
	if resolution < 2 {
		resolution = 2 // Need at least two points
	}

	lut := make([]Point, resolution)
	for i := range resolution {
		t := float64(i) / float64(resolution-1)

		// Standard cubic Bézier formula
		it := 1.0 - t
		it2 := it * it
		t2 := t * t

		x := it*it2*h.P0.X + 3*it2*t*h.P1.X + 3*it*t2*h.P2.X + t*t2*h.P3.X
		y := it*it2*h.P0.Y + 3*it2*t*h.P1.Y + 3*it*t2*h.P2.Y + t*t2*h.P3.Y

		lut[i] = Point{X: x, Y: y}
	}

	return &RasterizedBezier{points: lut}
}

// Sample finds the approximate Y for a given X by interpolating between pre-computed points.
func (ab *RasterizedBezier) Sample(x float64) float64 {
	// Handle edge cases where x is outside the bounds of our curve.
	if x <= ab.points[0].X {
		return ab.points[0].Y
	}
	lastPoint := ab.points[len(ab.points)-1]
	if x >= lastPoint.X {
		return lastPoint.Y
	}

	// Find the two points in our lookup table that the x value falls between.
	for i := 0; i < len(ab.points)-1; i++ {
		p1 := ab.points[i]
		p2 := ab.points[i+1]

		if x >= p1.X && x <= p2.X {
			// Found the segment. Now, linearly interpolate.

			// How far into the segment is our x? (a value between 0 and 1)
			t := (x - p1.X) / (p2.X - p1.X)

			// Apply that same proportion to the Y values.
			return p1.Y + t*(p2.Y-p1.Y)
		}
	}

	return lastPoint.Y // Should not be reached if x is within bounds
}
