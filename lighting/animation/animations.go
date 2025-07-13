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
					Point{X: 0, Y: 0.75},
					Point{X: 1, Y: 0.75},
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
				Offset:       -0.055 * 1,
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
				Offset:       -0.055 * 2,
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
				Offset:       -0.055 * 3,
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
				Offset:       -0.055 * 4,
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
				Offset:       -0.055 * 5,
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
				Offset:       -0.055 * 6,
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
				Offset:       -0.055 * 7,
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
				Offset:       -0.055 * 8,
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
				Offset:       -0.055 * 10,
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
				Offset:       -0.055 * 11,
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
				Offset:       -0.055 * 12,
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
				Offset:       -0.055 * 13,
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
				Offset:       -0.055 * 14,
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
				Offset:       -0.055 * 15,
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
				Offset:       -0.055 * 16,
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
				Offset:       -0.055 * 17,
			},
		},
		FPS:             60,
		DurationSeconds: 5,
	}
}
