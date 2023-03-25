package tracer

type HitRecord struct {
	P         Point3
	Normal    Vec3
	Material  Material
	T         float64
	FrontFace bool
}

func (h *HitRecord) SetFaceNormal(ray Ray, outwardNormal Vec3) {
	h.FrontFace = ray.Direction.Dot(outwardNormal) < 0
	if h.FrontFace {
		h.Normal = outwardNormal
	} else {
		h.Normal = outwardNormal.Negate()
	}
}

type Hittable interface {
	Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord)
}
