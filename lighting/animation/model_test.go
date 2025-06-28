package animation

import (
	"testing"
)

func TestCurve_Sample(t *testing.T) {
	tests := []struct {
		name string
		c    Curve
	}{
		{
			name: "test 1",
			c: Curve{
				{
					P0: Point{X: 0, Y: 0},
					P1: Point{X: 0.25, Y: 0},
					P2: Point{X: 0.75, Y: 1.0},
					P3: Point{X: 1, Y: 1},
				},
				{
					P0: Point{X: 1, Y: 1},
					P1: Point{X: 1.25, Y: 1.0},
					P2: Point{X: 1.75, Y: 0},
					P3: Point{X: 2, Y: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			xs := []float64{}
			ys := []float64{}
			bytes := []byte{}

			for i := range 41 {
				x := float64(i) / 20.0
				y := tt.c.Sample(x)
				b := tt.c.SampleByte(x)
				xs = append(xs, x)
				ys = append(ys, y)
				bytes = append(bytes, b)
			}

			t.Log("xs", xs)
			t.Log("ys", ys)
			t.Log("bytes", bytes)
		})
	}
}
