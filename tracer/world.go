package tracer

type World struct {
	Objects []Hittable
}

func CreateWorld() World {
	return World{
		Objects: []Hittable{},
	}
}

func (h *World) Add(obj Hittable) {
	h.Objects = append(h.Objects, obj)
}

func (h *World) Clear() {
	h.Objects = []Hittable{}
}

func (h *World) Hit(ray Ray, tMin float64, tMax float64) (bool, HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	rec := HitRecord{}

	for _, obj := range h.Objects {
		hit, tempRec := obj.Hit(ray, tMin, closestSoFar)

		if hit {
			hitAnything = true
			closestSoFar = tempRec.T
			rec = tempRec
		}
	}

	return hitAnything, rec
}
