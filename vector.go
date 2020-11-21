package main

type Vec2i struct {
	X, Y int
}

func (vi Vec2i) Vec2f() Vec2f {
	return Vec2f{
		X: float64(vi.X),
		Y: float64(vi.Y),
	}
}

type Vec2f struct {
	X, Y float64
}

func (vf Vec2f) Vec2i() Vec2i {
	return Vec2i{
		X: int(vf.X),
		Y: int(vf.Y),
	}
}
