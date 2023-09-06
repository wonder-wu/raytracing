package rt

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(rIn Ray, hitRecord *HitRecord, rd *rand.Rand) (bool, Vector3, Ray)
}

type Lambertian struct {
	Albedo Vector3
}

func NewLamertian(a Vector3) *Lambertian {
	return &Lambertian{Albedo: a}
}

func (l *Lambertian) Scatter(rIn Ray, hitRecord *HitRecord, rd *rand.Rand) (b bool, attenuation Vector3, scattered Ray) {
	scatterDir := hitRecord.Normal.Add(Vec3RandomNewInUnitCircle(rd))
	scattered = NewRay(hitRecord.P, scatterDir)
	attenuation = l.Albedo
	b = true
	return
}

type Metal struct {
	Albedo Vector3
	Fuzz   float64
}

func NewMetal(a Vector3, fuzz float64) *Metal {
	if fuzz < 0.0 {
		fuzz = 1.0
	}
	return &Metal{Albedo: a, Fuzz: fuzz}
}

func reflect(uv, n Vector3) Vector3 {
	return uv.Add(Vec3MulFloat(n, -2.0*Vec3Dot(uv, n)))
}

func (m *Metal) Scatter(rIn Ray, hitRecord *HitRecord, rd *rand.Rand) (b bool, attenuation Vector3, scattered Ray) {
	rInUnit := rIn.Dir.Normalize()
	reflectDir := reflect(rInUnit, hitRecord.Normal)
	scattered = NewRay(hitRecord.P, reflectDir.Add(Vec3MulFloat(Vec3RandomNewInUnitSphere(rd), m.Fuzz)))
	attenuation = m.Albedo
	b = Vec3Dot(scattered.Dir, hitRecord.Normal) > 0
	return
}

func refract(uv, n Vector3, etaiOverEtat float64) Vector3 {
	cosTheta := math.Min(Vec3Dot(uv.MulFloat(-1.0), n), 1.0)
	rOutParallel := Vec3Add(uv, n.MulFloat(cosTheta)).MulFloat(etaiOverEtat)
	rOutPerpendicular := n.MulFloat(-1.0).MulFloat(math.Sqrt(1.0 - rOutParallel.LengthSquared()))
	return Vec3Add(rOutParallel, rOutPerpendicular)
}

func schlick(cosine, refIdx float64) float64 {

	r0 := (1.0 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1.0-r0)*math.Pow(1.0-cosine, 5.0)
}

type Dielectric struct {
	RefIdx float64
}

func NewDielectric(refidx float64) *Dielectric {
	return &Dielectric{RefIdx: refidx}
}

func (r *Dielectric) Scatter(rIn Ray, hitRecord *HitRecord, rd *rand.Rand) (b bool, attenuation Vector3, scattered Ray) {
	b = true
	attenuation = Vec3(1.0, 1.0, 1.0)
	var etaiOverEtat float64
	if hitRecord.FrontFace {
		etaiOverEtat = 1.0 / r.RefIdx
	} else {
		etaiOverEtat = r.RefIdx
	}

	uintDir := rIn.Dir.Normalize()

	cosTheta := math.Min(Vec3Dot(uintDir.MulFloat(-1.0), hitRecord.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	if etaiOverEtat*sinTheta > 1.0 {
		//reflect
		reflected := reflect(uintDir, hitRecord.Normal)
		scattered = NewRay(hitRecord.P, reflected)
		return
	}
	reflectProb := schlick(cosTheta, etaiOverEtat)
	if RandomNew(rd) < reflectProb {
		reflected := reflect(uintDir, hitRecord.Normal)
		scattered = NewRay(hitRecord.P, reflected)

		return

	}
	refracted := refract(uintDir, hitRecord.Normal, etaiOverEtat)
	scattered = NewRay(hitRecord.P, refracted)

	return
}
