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
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       0,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 1,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 2,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 3,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 4,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 5,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 6,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 7,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 8,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 10,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 11,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 12,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 13,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 14,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 15,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 16,
			},
			{

				Value1: Curve{
					Point{X: 0, Y: 0},
					Point{X: 1, Y: 1},
				},
				Value2: Curve{
					Point{X: 0, Y: 1},
					Point{X: 1, Y: 1},
				},
				Value3: Curve{
					Point{X: 0, Y: 0},
					Point{X: 0.25, Y: 1},
					Point{X: 0.5, Y: 0},
				},
				Brightness: Curve{
					Point{X: 0, Y: 0.1},
					Point{X: 1, Y: 0.1},
				},
				Interpolator: Linear,
				ValueType:    HSV,
				Offset:       -0.025 * 17,
			},
		},
		FPS:             60,
		DurationSeconds: 1,
	}
}
