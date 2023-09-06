package rt

import (
	"math"
)

type Sphere struct {
	Center Vector3
	Radius float64
	Mat    Material
}

func NewSphere(center Vector3, rad float64, mat Material) *Sphere {
	return &Sphere{center, rad, mat}
}

func (s *Sphere) Hit(r Ray, tMin, tMax float64, hitRecord *HitRecord) bool {
	oc := r.Ori.Sub(s.Center)
	a := r.Dir.LengthSquared()
	b := Vec3Dot(oc, r.Dir)
	c := oc.LengthSquared() - s.Radius*s.Radius
	solves := b*b - a*c
	if solves > 0 {
		root := math.Sqrt(solves)
		temp := (-b - root) / a
		if temp < tMax && temp > tMin {
			hitRecord.T = temp
			hitRecord.P = r.At(temp)
			outwardNormal := hitRecord.P.Sub(s.Center).DivFloat(s.Radius)
			hitRecord.SetFaceNoraml(r, outwardNormal)
			hitRecord.Mat = s.Mat
			return true
		}

		temp = (-b + root) / a
		if temp < tMax && temp > tMin {
			hitRecord.T = temp
			hitRecord.P = r.At(temp)
			outwardNormal := hitRecord.P.Sub(s.Center).DivFloat(s.Radius)
			hitRecord.SetFaceNoraml(r, outwardNormal)
			hitRecord.Mat = s.Mat
			return true
		}

	}
	return false

}

func (s *Sphere) BoundingBox() AABB {
	return NewAABB(s.Center.Sub(Vec3(1.0, 1.0, 1.0).MulFloat(s.Radius)), s.Center.Add(Vec3(1.0, 1.0, 1.0).MulFloat(s.Radius)))

}
