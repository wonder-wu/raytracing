package rt

import (
	"math"
	"math/rand"
)

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

//Vec3 create a new vector3
func Vec3(param ...float64) Vector3 {

	switch len(param) {
	case 0:
		return Vector3{0, 0, 0}
	case 1:
		return Vector3{param[0], param[0], param[0]}
	case 2:
		return Vector3{param[0], param[1], param[1]}
	default:
		return Vector3{param[0], param[1], param[2]}
	}
}

func Vec3Sub(a, b Vector3) Vector3 {
	return Vector3{
		a.X - b.X,
		a.Y - b.Y,
		a.Z - b.Z}
}

func Vec3Add(a, b Vector3) Vector3 {
	return Vector3{
		a.X + b.X,
		a.Y + b.Y,
		a.Z + b.Z}
}

func Vec3MulFloat(a Vector3, b float64) Vector3 {
	return Vector3{
		a.X * b,
		a.Y * b,
		a.Z * b}
}

func Vec3DivFloat(a Vector3, b float64) Vector3 {
	return Vec3MulFloat(a, 1.0/b)
}

func Vec3MulVec3(a, b Vector3) Vector3 {
	return Vector3{
		a.X * b.X,
		a.Y * b.Y,
		a.Z * b.Z}
}

func Vec3Dot(a, b Vector3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Vec3Cross(a, b Vector3) Vector3 {
	return Vector3{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}

func Vec3Normalize(a Vector3) Vector3 {
	return Vec3DivFloat(a, a.Length())
}

func Vec3Random() Vector3 {
	return Vec3(Random(), Random(), Random())
}
func Vec3RandomBetween(min, max float64) Vector3 {
	return Vec3(RandomBetween(min, max), RandomBetween(min, max), RandomBetween(min, max))
}
func Vec3RandomInUnitSphere() Vector3 {
	for {
		p := Vec3RandomBetween(-1.0, 1.0)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}

}
func Vec3RandomInUnitCircle() Vector3 {

	a := RandomBetween(0.0, 2.0) * PI
	z := RandomBetween(-1.0, 1.0)
	r := math.Sqrt(1.0 - z*z)

	return Vec3(r*math.Cos(a), r*math.Sin(a), z)

}

func Vec3RandomNew(r *rand.Rand) Vector3 {
	return Vec3(RandomNew(r), RandomNew(r), RandomNew(r))
}
func Vec3RandomBetweenNew(r *rand.Rand, min, max float64) Vector3 {
	return Vec3(RandomNewBetween(r, min, max), RandomNewBetween(r, min, max), RandomNewBetween(r, min, max))
}
func Vec3RandomNewInUnitSphere(r *rand.Rand) Vector3 {
	for {
		p := Vec3RandomBetweenNew(r, -1.0, 1.0)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}

}
func Vec3RandomNewInUnitDisk(rd *rand.Rand) Vector3 {
	for {
		p := Vec3RandomBetweenNew(rd, -1.0, 1.0)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func Vec3RandomNewInUnitCircle(rd *rand.Rand) Vector3 {

	a := RandomNewBetween(rd, 0.0, 2.0) * PI
	z := RandomNewBetween(rd, -1.0, 1.0)
	r := math.Sqrt(1.0 - z*z)

	return Vec3(r*math.Cos(a), r*math.Sin(a), z)

}

func Length(v Vector3) float64 {
	return v.Length()
}
func (v Vector3) Add(b Vector3) Vector3 {
	return Vec3Add(v, b)
}
func (v Vector3) Sub(b Vector3) Vector3 {
	return Vec3Sub(v, b)
}
func (v Vector3) MulFloat(f float64) Vector3 {
	return Vec3MulFloat(v, f)
}
func (v Vector3) DivFloat(f float64) Vector3 {
	return Vec3DivFloat(v, f)
}
func (v Vector3) MulVec3(b Vector3) Vector3 {
	return Vec3MulVec3(v, b)
}
func (v Vector3) Dot(b Vector3) float64 {
	return Vec3Dot(v, b)
}
func (v Vector3) Cross(b Vector3) Vector3 {
	return Vec3Cross(v, b)
}
func (v Vector3) Normalize() Vector3 {
	return Vec3Normalize(v)
}
func (v Vector3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vector3) Get(i int) float64 {
	switch {
	case i <= 0:
		return v.X
	case i == 1:
		return v.Y
	case i >= 2:
		return v.Z
	default:
		return 0
	}
}
