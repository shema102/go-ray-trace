package renderer

import (
	"math"
	"rt/util"
)

type Vec3 struct {
	X, Y, Z float64
}

func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec3) Mul(v2 Vec3) Vec3 {
	return Vec3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vec3) MulScalar(s float64) Vec3 {
	return Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3) DivScalar(s float64) Vec3 {
	return Vec3{v.X / s, v.Y / s, v.Z / s}
}

func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.Y*v2.Z - v.Z*v2.Y,
		v.Z*v2.X - v.X*v2.Z,
		v.X*v2.Y - v.Y*v2.X,
	}
}

func (v Vec3) Negate() Vec3 {
	return Vec3{-v.X, -v.Y, -v.Z}
}

func (v Vec3) UnitVector() Vec3 {
	return v.DivScalar(v.Length())
}

func (v Vec3) Reflect(n Vec3) Vec3 {
	return v.Sub(n.MulScalar(2 * v.Dot(n)))
}

//func (v Vec3) Refract(n Vec3, etaiOverEtat float64) (Vec3, bool) {
//	uv := v.UnitVector()
//	dt := uv.Dot(n)
//	discriminant := 1.0 - etaiOverEtat*etaiOverEtat*(1-dt*dt)
//	if discriminant > 0 {
//		refracted := uv.Sub(n.MulScalar(dt)).MulScalar(etaiOverEtat).Sub(n.MulScalar(math.Sqrt(discriminant)))
//		return refracted, true
//	} else {
//		return Vec3{}, false
//	}
//}

func (v Vec3) Refract(n Vec3, etaiOverEtat float64) Vec3 {
	cosTheta := math.Min(n.Dot(v.Negate()), 1.0)
	rOutPerp := v.Add(n.MulScalar(cosTheta)).MulScalar(etaiOverEtat)
	rOutParallel := n.MulScalar(-math.Sqrt(math.Abs(1.0 - rOutPerp.LengthSquared())))
	refracted := rOutPerp.Add(rOutParallel)
	return refracted
}

func (v Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vec3) NearZero() bool {
	s := 1e-8
	return (math.Abs(v.X) < s) && (math.Abs(v.Y) < s) && (math.Abs(v.Z) < s)
}

type Color = Vec3

type Point3 = Vec3

func RandomVec3() Vec3 {
	return Vec3{util.Random(), util.Random(), util.Random()}
}

func RandomVec3Range(min, max float64) Vec3 {
	return Vec3{util.RandomRange(min, max), util.RandomRange(min, max), util.RandomRange(min, max)}
}

func RandomInUnitSphere() Vec3 {
	for {
		p := RandomVec3Range(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func RandomUnitVector() Vec3 {
	return RandomInUnitSphere().UnitVector()
}

func RandomInHemisphere(normal Vec3) Vec3 {
	inUnitSphere := RandomInUnitSphere()
	if inUnitSphere.Dot(normal) > 0.0 {
		return inUnitSphere
	}
	return inUnitSphere.Negate()
}
