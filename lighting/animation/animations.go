package animation

func AnimationPulse() Animation {
	return Animation{
		Name: "pulse",
		Tracks: []Track{
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 0.75},
					Point{X: 1, Y: 0.75},
				},
				Value3: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 0.75},
					Point{X: 1, Y: 0.75},
				},
				Value3: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
			},
		},
		FPS:             60,
		DurationSeconds: 5,
	}
}
