package renderer

import (
	"math"
	"rt/util"
)

type Material interface {
	Scatter(rIn Ray, rec HitRecord) (bool, Ray, Vec3)
}

type Lambertian struct {
	Albedo Vec3
}

func (l Lambertian) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Vec3) {
	scatterDirection := rec.Normal.Add(RandomUnitVector())

	if scatterDirection.NearZero() {
		scatterDirection = rec.Normal
	}

	scattered := Ray{rec.P, scatterDirection}
	attenuation := l.Albedo

	return true, scattered, attenuation
}

type Metal struct {
	Albedo Vec3
	Fuzz   float64
}

func (m Metal) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Vec3) {
	reflected := rIn.Direction.UnitVector().Reflect(rec.Normal)
	scattered := Ray{rec.P, reflected.Add(RandomInUnitSphere().MulScalar(m.Fuzz))}

	attenuation := m.Albedo

	return scattered.Direction.Dot(rec.Normal) > 0, scattered, attenuation
}

type Dielectric struct {
	RefractiveIndex float64
}

func (d Dielectric) Schlick(cosine, refractionRatio float64) float64 {
	r0 := (1 - refractionRatio) / (1 + refractionRatio)
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

func (d Dielectric) Scatter(rIn Ray, rec HitRecord) (bool, Ray, Vec3) {
	attenuation := Vec3{1.0, 1.0, 1.0}

	refractionRatio := 0.0

	if rec.FrontFace {
		refractionRatio = 1.0 / d.RefractiveIndex
	} else {
		refractionRatio = d.RefractiveIndex
	}

	cosTheta := math.Min(rIn.Direction.Negate().Dot(rec.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0
	direction := Vec3{}

	if cannotRefract || d.Schlick(cosTheta, refractionRatio) > util.Random() {
		direction = rIn.Direction.Reflect(rec.Normal)
	} else {
		direction = rIn.Direction.Refract(rec.Normal, refractionRatio)
	}

	scattered := Ray{rec.P, direction}

	return true, scattered, attenuation
}
