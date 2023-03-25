package tracer

import (
	"math"
	"rt/util"
)

type Camera struct {
	Origin, LowerLeftCorner, Horizontal, Vertical Vec3
	ViewPortHeight, ViewPortWidth                 float64
}

func NewCamera(vertFov, aspectRatio float64) Camera {
	theta := util.DegreesToRadians(vertFov)
	halfHeight := math.Tan(theta / 2)

	viewPortHeight := 2.0 * halfHeight
	viewPortWidth := aspectRatio * viewPortHeight

	focalLength := 1.0

	origin := Vec3{}
	horizontal := Vec3{X: viewPortWidth}
	vertical := Vec3{Y: viewPortHeight}
	lowerLeftCorner := origin.Sub(horizontal.DivScalar(2)).Sub(vertical.DivScalar(2)).Sub(Vec3{Z: focalLength})

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
