package rt

type Ray struct {
	Ori Vector3
	Dir Vector3
}

func NewRay(origin Vector3, direction Vector3) Ray {
	return Ray{origin, direction}
}
func (r Ray) Origin() Vector3 {
	return r.Ori
}
func (r Ray) Direction() Vector3 {
	return r.Dir
}

func (r Ray) At(t float64) Vector3 {
	return Vec3Add(r.Ori, Vec3MulFloat(r.Dir, t))
}
