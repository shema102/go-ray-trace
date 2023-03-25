package tracer

type Plane struct {
	Normal   Vec3
	Distance float64
	Material Material
}

func (p Plane) Hit(r Ray, tMin float64, tMax float64) (bool, HitRecord) {
	denominator := r.Direction.Dot(p.Normal)

	// if ray is parallel to plane
	if denominator == 0 {
		return false, HitRecord{}
	}

	t := (p.Distance - r.Origin.Dot(p.Normal)) / denominator

	if t < tMin || tMax < t {
		return false, HitRecord{}
	}

	rec := HitRecord{}

	rec.T = t
	rec.P = r.At(rec.T)
	rec.Normal = p.Normal
	rec.SetFaceNormal(r, rec.Normal)
	rec.Material = p.Material

	return true, rec
}
