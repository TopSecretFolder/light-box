package animation

func AnimationPulse() Animation {
	return Animation{
		Name: "pulse",
		Red: Curve{
			{
				P0: Point{X: 0, Y: 1},
				P1: Point{X: 0, Y: 1},
				P2: Point{X: 2, Y: 1},
				P3: Point{X: 2, Y: 1},
			},
		},
		Green: Curve{
			{
				P0: Point{X: 0, Y: 0},
				P1: Point{X: 0, Y: 0},
				P2: Point{X: 2, Y: 0},
				P3: Point{X: 2, Y: 0},
			},
		},
		Blue: Curve{
			{
				P0: Point{X: 0, Y: 0},
				P1: Point{X: 0, Y: 0},
				P2: Point{X: 2, Y: 0},
				P3: Point{X: 2, Y: 0},
			},
		},
		Brightness: Curve{
			{
				P0: Point{X: 0, Y: 0},
				P1: Point{X: 0.25, Y: 0},
				P2: Point{X: 0.75, Y: 0.1},
				P3: Point{X: 1, Y: 0.1},
			},
			{
				P0: Point{X: 1, Y: 0.1},
				P1: Point{X: 1.25, Y: 0.1},
				P2: Point{X: 1.75, Y: 0},
				P3: Point{X: 2, Y: 0},
			},
		},
		Frames:          120,
		DurationSeconds: 10,
	}
}
