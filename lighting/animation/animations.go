package animation

func AnimationPulse() Animation {
	return Animation{
		Name: "pulse",
		Value1: Curve{
			Point{X: 0, Y: 0},
			Point{X: 1, Y: 1},
		},
		Value2: Curve{
			Point{X: 0, Y: 1},
			Point{X: 1, Y: 1},
		},
		Value3: Curve{
			Point{X: 0, Y: 1},
			Point{X: 1, Y: 1},
		},
		Brightness: Curve{
			Point{X: 0, Y: 1},
			Point{X: 0.5, Y: 0},
			Point{X: 1, Y: 1},
		},
		Frames:          120,
		DurationSeconds: 10,
		Interpolator:    Cubic,
		ValueType:       HSV,
	}
}
