package animation

import (
	"math"

	"gonum.org/v1/gonum/interp"
)

type Curve []Point

func (c Curve) Sample(head float64, interpolator Interpolator) float64 {
	switch interpolator {
	case Linear:
		return LinearInterpolator(c)(head)
	case Cubic:
		return CubicInterpolator(c)(head)
	}
	return 0
}

func (c Curve) SampleByte(head float64, interpolator Interpolator) byte {
	y := c.Sample(head, interpolator)
	return byte(math.Floor(y * 255))

}

func (c Curve) SampleByteMaxValue(head float64, maxValue int, interpolator Interpolator) byte {
	y := c.Sample(head, interpolator)

	return byte(math.Floor(y * float64(maxValue)))

}

func (c Curve) Domain() float64 {
	if len(c) == 0 {
		return 0
	}
	return c[len(c)-1].X - c[0].X
}

type Track struct {
	Interpolator Interpolator `json:"interpolator"`
	ValueType    ValueType    `json:"value_type"`
	Value1       Curve        `json:"value_1"` // r or h
	Value2       Curve        `json:"value_2"` // g or s
	Value3       Curve        `json:"value_3"` // b or v
	Brightness   Curve        `json:"brightness"`
}

type Interpolator string

const (
	Linear Interpolator = "linear"
	Cubic  Interpolator = "cubic"
)

type ValueType string

const (
	RGB ValueType = "rgb"
	HSV ValueType = "hsv"
)

type Animation struct {
	Name            string  `json:"name"`
	Tracks          []Track `json:"tracks"`
	FPS             int     `json:"fps"`
	DurationSeconds float64 `json:"duration_seconds"`
}

func (a Animation) GetBrightness(head float64, trackIndex int) byte {
	return a.track(trackIndex).Brightness.SampleByteMaxValue(head, 10, a.track(trackIndex).Interpolator)
}

func (a Animation) Sample(head float64, trackIndex int) (r byte, g byte, b byte) {
	v1 := a.track(trackIndex).Value1.Sample(head, a.track(trackIndex).Interpolator)
	v2 := a.track(trackIndex).Value2.Sample(head, a.track(trackIndex).Interpolator)
	v3 := a.track(trackIndex).Value3.Sample(head, a.track(trackIndex).Interpolator)

	if a.track(trackIndex).ValueType == HSV {
		r, g, b := HSVToRGB(v1, v2, v3)
		return ToByte(r, g, b)
	}
	return ToByte(v1, v2, v3)
}

func (a Animation) track(i int) Track {
	if i < len(a.Tracks) {
		return a.Tracks[i]
	}

	// by default if no track exists for index i, then return zero valued
	// track
	return Track{
		Interpolator: Linear,
		ValueType:    HSV,
		Value1: Curve{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 1,
				Y: 0,
			},
		},
		Value2: Curve{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 1,
				Y: 0,
			},
		},
		Value3: Curve{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 1,
				Y: 0,
			},
		},
		Brightness: Curve{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 1,
				Y: 0,
			},
		},
	}
}

func ToByte(v1 float64, v2 float64, v3 float64) (byte, byte, byte) {
	return byte(math.Floor(v1 * 255)), byte(math.Floor(v2 * 255)), byte(math.Floor(v3 * 255))
}

// HSVToRGB converts an HSV color value to RGB.
// Conversion formula adapted from http://en.wikipedia.org/wiki/HSV_color_space.
// Assumes h, s, and v are contained in the set [0, 1] and
// returns r, g, and b in the set [0, 1].
func HSVToRGB(h, s, v float64) (float64, float64, float64) {
	if s == 0 {
		return v, v, v
	}

	h *= 6
	i := math.Floor(h)
	f := h - i

	p := v * (1 - s)
	q := v * (1 - s*f)
	t := v * (1 - s*(1-f))

	switch i {
	case 0:
		return v, t, p
	case 1:
		return q, v, p
	case 2:
		return p, v, t
	case 3:
		return p, q, v
	case 4:
		return t, p, v
	default: // case 5:
		return v, p, q
	}
}

func (a Animation) Domain(trackIndex int) float64 {
	d1 := a.track(trackIndex).Value1.Domain()
	d2 := a.track(trackIndex).Value2.Domain()
	d3 := a.track(trackIndex).Value3.Domain()
	dbr := a.track(trackIndex).Brightness.Domain()
	return math.Max(math.Max(math.Max(d1, d2), d3), dbr)
}

// Point represents a 2D point.
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func LinearInterpolator(points Curve) func(t float64) float64 {
	if len(points) < 2 {
		// Cannot interpolate with less than 2 points.
		// Return a function that returns NaN or handles the error as you see fit.
		return func(t float64) float64 { return 0 }
	}

	// Separate the points into x and y slices for the library.
	xs := make([]float64, len(points))
	ys := make([]float64, len(points))
	for i, p := range points {
		xs[i] = p.X
		ys[i] = p.Y
	}

	// Create a new linear interpolator from the gonum library.
	linear := interp.PiecewiseLinear{}
	linear.Fit(xs, ys)

	// Return a closure that can be called with any t.
	return func(t float64) float64 {
		return linear.Predict(t)
	}
}

func CubicInterpolator(points Curve) func(t float64) float64 {
	if len(points) < 2 {
		return func(t float64) float64 { return 0 }
	}

	xs := make([]float64, len(points))
	ys := make([]float64, len(points))
	for i, p := range points {
		xs[i] = p.X
		ys[i] = p.Y
	}

	// Create a new natural cubic spline interpolator.
	// Natural splines have zero second derivatives at the endpoints,
	// which is a common and safe choice.
	spline := interp.ClampedCubic{}
	spline.Fit(xs, ys)

	// Return the prediction function.
	return func(t float64) float64 {
		return spline.Predict(t)
	}
}
