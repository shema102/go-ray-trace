package renderer

import (
	"math"
)

type HitRecord struct {
	P         Point3
	Normal    Vec3
	Material  Material
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(r Ray, outwardNormal Vec3) {
	h.FrontFace = r.Direction.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negate()
	}
}

type Hittable interface {
	Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord)
}

type Sphere struct {
	Center   Point3
	Radius   float64
	Material Material
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
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
	rec.P = r.At(rec.T)
	outwardNormal := rec.P.Sub(s.Center).DivScalar(s.Radius)
	rec.SetFaceNormal(r, outwardNormal)
	rec.Material = s.Material

	return true, rec
}

type HittableList struct {
	Objects []Hittable
}

func CreateHittableList() HittableList {
	return HittableList{
		Objects: []Hittable{},
	}
}

func (h *HittableList) Add(obj Hittable) {
	h.Objects = append(h.Objects, obj)
}

func (h *HittableList) Clear() {
	h.Objects = []Hittable{}
}

func (h *HittableList) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	rec := HitRecord{}

	for _, obj := range h.Objects {
		hit, tempRec := obj.Hit(r, tMin, closestSoFar)

		if hit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}

	return hitAnything, rec
}
