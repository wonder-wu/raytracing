package rt

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkRandomWithNewRand(b *testing.B) {
	source := rand.NewSource(time.Now().Unix())
	rd := rand.New(source)
	for i := 0; i < b.N; i++ {
		Vec3RandomNewInUnitCircle(rd)
	}

}

func BenchmarkRandomWithoutNewRand(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Vec3RandomInUnitCircle()
	}
}

func BenchmarkVectorGet(b *testing.B) {

	v := Vec3(1.0, 2.0, 3.0)
	for i := 0; i < b.N; i++ {
		_ = v.Get(0)
		_ = v.Get(1)
		_ = math.Sqrt(v.Get(2)) * math.Pow(3.66, 100.66)
	}
}

func BenchmarkVectorGetRaw(b *testing.B) {

	v := Vec3(1.0, 2.0, 3.0)
	for i := 0; i < b.N; i++ {

		_ = math.Sqrt(v.X) * math.Pow(3.66, 100.66)
		_ = v.Y
		_ = v.Z
	}
}
