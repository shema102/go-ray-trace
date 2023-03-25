package tracer

import (
	"math"
	"rt/util"
)

type Camera struct {
	Origin, LowerLeftCorner, Horizontal, Vertical Vec3
	ViewPortHeight, ViewPortWidth                 float64
}

func NewCamera(lookFrom, lookAt Point3, upVector Vec3, vertFov, aspectRatio float64) Camera {
	theta := util.DegreesToRadians(vertFov)
	halfHeight := math.Tan(theta / 2)

	viewPortHeight := 2.0 * halfHeight
	viewPortWidth := aspectRatio * viewPortHeight

	focalLength := 1.0

	origin := lookFrom
	w := lookFrom.Sub(lookAt).UnitVector()
	u := upVector.Cross(w).UnitVector()
	v := w.Cross(u)

	horizontal := u.MulScalar(viewPortWidth)
	vertical := v.MulScalar(viewPortHeight)
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(w.MulScalar(focalLength))

	return Camera{
		Origin:          origin,
		LowerLeftCorner: lowerLeftCorner,
		Horizontal:      horizontal,
		Vertical:        vertical,
		ViewPortHeight:  viewPortHeight,
		ViewPortWidth:   viewPortWidth,
	}
}

func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		Origin:    c.Origin,
		Direction: c.LowerLeftCorner.Add(c.Horizontal.MulScalar(u)).Add(c.Vertical.MulScalar(v)).Sub(c.Origin),
	}
}
