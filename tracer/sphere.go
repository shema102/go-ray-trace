package tracer

import (
	"math"
)

type Sphere struct {
	Center   Point3
	Radius   float64
	Material Material
}

func (s Sphere) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := ray.Origin.Sub(s.Center)
	a := ray.Direction.LengthSquared()
	halfB := oc.Dot(ray.Direction)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c

	if discriminant < 0 {
		return false, HitRecord{}
	}

	sqrt := math.Sqrt(discriminant)

	root := (-halfB - sqrt) / a

	if root < tMin || tMax < root {
		root = (-halfB + sqrt) / a

		if root < tMin || tMax < root {
			return false, HitRecord{}
		}
	}

	rec := HitRecord{}

	rec.T = root
	rec.P = ray.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).DivScalar(s.Radius)
	rec.SetFaceNormal(ray, outwardNormal)
	rec.Material = s.Material

	return true, rec
}
