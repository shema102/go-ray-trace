package tracer

import (
	"math"
)

type Ray struct {
	Origin, Direction Vec3
}

func (r Ray) At(t float64) Vec3 {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

func TraceRay(r Ray, world *World, depth int) Color {
	hit, rec := world.Hit(r, 0.001, math.MaxFloat64)

	if depth <= 0 {
		return Color{0, 0, 0}
	}

	if hit {
		isScattered, scatteredRay, attenuation := rec.Material.Scatter(r, rec)

		if isScattered {
			return attenuation.Mul(TraceRay(scatteredRay, world, depth-1))
		}

		return Color{0, 0, 0}
	}

	unitDirection := r.Direction.UnitVector()

	t := 0.5 * (unitDirection.Y + 1.0)

	return Color{1.0, 1.0, 1.0}.MulScalar(1.0 - t).Add(Color{0.5, 0.7, 1.0}.MulScalar(t))
}
