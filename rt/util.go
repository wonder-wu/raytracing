package rt

import (
	"image/color"
	"math"
	"math/rand"
)

const (
	INFINTE float64 = math.MaxFloat64
	PI      float64 = 3.1415926535897932385
)

func DegreeToRadius(deg float64) float64 {
	return deg * PI / 180.0
}

func SampleColorToRGBA(col Vector3, samplesPerPixel int) color.RGBA {
	r := col.X
	g := col.Y
	b := col.Z

	scale := 1.0 / float64(samplesPerPixel)
	// r *= scale
	// g *= scale
	// b *= scale
	r = math.Sqrt(r * scale)
	g = math.Sqrt(g * scale)
	b = math.Sqrt(b * scale)

	return color.RGBA{
		uint8(Clamp(r, 0.0, 1.0) * 255),
		uint8(Clamp(g, 0.0, 1.0) * 255),
		uint8(Clamp(b, 0.0, 1.0) * 255),
		255,
	}

}

func NormalVec3ToRGBA(v Vector3) color.RGBA {
	return color.RGBA{
		uint8(v.X * 255),
		uint8(v.Y * 255),
		uint8(v.Z * 255),
		255,
	}
}

func RandomNew(r *rand.Rand) float64 {
	return r.Float64()
}
func RandomNewBetween(r *rand.Rand, min, max float64) float64 {
	return min + (max-min)*RandomNew(r)
}

func Random() float64 {

	return rand.Float64()
}

func RandomBetween(min, max float64) float64 {
	return min + (max-min)*Random()
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
