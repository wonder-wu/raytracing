package rt

import (
	"math"
)

type AABB struct {
	min Vector3
	max Vector3
}

func NewAABB(a, b Vector3) AABB {
	return AABB{a, b}
}

func (aabb AABB) Hit(r Ray, tMin, tMax float64) bool {
	for i := 0; i < 3; i++ {
		// r1 := (aabb.min.Get(i) - r.Ori.Get(i)) / r.Dir.Get(i)
		// r2 := (aabb.max.Get(i) - r.Ori.Get(i)) / r.Dir.Get(i)
		// t0 := math.Min(r1, r2)

		// t1 := math.Max(r1, r2)

		// tMin = math.Max(t0, tMin)
		// tMax = math.Min(t1, tMax)
		// if tMax <= tMin {
		// 	return false
		// }
		oi := r.Ori.Get(i)
		amini := aabb.min.Get(i)
		amaxi := aabb.max.Get(i)
		invD := 1.0 / r.Dir.Get(i)
		t0 := (amini - oi) * invD
		t1 := (amaxi - oi) * invD
		if invD < 0.0 {
			t0, t1 = t1, t0
		}
		if t0 > tMin {
			tMin = t0
		}
		if t1 < tMax {
			tMax = t1
		}
		if tMax <= tMin {
			return false
		}

	}
	return true
}

func SurroundingBox(box1, box2 AABB) AABB {
	small := Vec3(math.Min(box1.min.X, box2.min.X),
		math.Min(box1.min.Y, box2.min.Y),
		math.Min(box1.min.Z, box2.min.Z))

	big := Vec3(math.Max(box1.max.X, box2.max.X),
		math.Max(box1.max.Y, box2.max.Y),
		math.Max(box1.max.Z, box2.max.Z))

	return AABB{small, big}

}

func BoxCompare(a, b AABB, axis int) bool {
	return a.min.Get(axis) < b.min.Get(axis)
}
