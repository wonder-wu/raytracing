package rt

import (
	"math"
	"math/rand"
)

type Camera struct {
	origin          Vector3
	lowerLeftCorner Vector3
	horizontal      Vector3
	vertical        Vector3
	u, v, w         Vector3
	lensRadius      float64
}

func NewCamera(lookFrom, lookAt, vUp Vector3, vfov, aspectRatio float64, aperture, focusDist float64) Camera {
	theta := DegreeToRadius(vfov)
	h := math.Tan(theta / 2.0)
	viewerPortHeight := 2.0 * h
	viewerPortwidth := viewerPortHeight * aspectRatio

	w := (lookFrom.Sub(lookAt)).Normalize()
	u := Vec3Cross(vUp, w).Normalize()
	v := Vec3Cross(w, u)

	origin := lookFrom

	horizontal := u.MulFloat(viewerPortwidth).MulFloat(focusDist)
	vertical := v.MulFloat(viewerPortHeight).MulFloat(focusDist)

	lowerLeftCorner := origin.Sub(horizontal.DivFloat(2.0)).Sub(vertical.DivFloat(2.0)).Sub(w.MulFloat(focusDist))

	return Camera{
		origin:          origin,
		lowerLeftCorner: lowerLeftCorner,
		horizontal:      horizontal,
		vertical:        vertical,
		u:               u,
		v:               v,
		w:               w,
		lensRadius:      aperture / 2.0,
	}

}

func (c *Camera) GetRay(u, v float64, rd *rand.Rand) Ray {
	random := Vec3RandomNewInUnitDisk(rd).MulFloat(c.lensRadius)
	offset := Vec3Add(c.u.MulFloat(random.X), c.v.MulFloat(random.Y))

	return NewRay(c.origin,
		c.lowerLeftCorner.Add(c.horizontal.MulFloat(u)).Add(c.vertical.MulFloat(v)).Sub(c.origin).Sub(offset))
}
