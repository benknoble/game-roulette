package roulette

import (
	"math/rand"
	"testing"
)

func TestRandSeedTen(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	// first ten values of Float64
	ten := []float64{
		0.6046602879796196,
		0.9405090880450124,
		0.6645600532184904,
		0.4377141871869802,
		0.4246374970712657,
		0.6868230728671094,
		0.06563701921747622,
		0.15651925473279124,
		0.09696951891448456,
		0.30091186058528707,
	}
	for i := range ten {
		test := r.Float64()
		if test != ten[i] {
			t.Error("New Random seeded with 1 gave (%v), expected (%v)", test, ten[i])
		}
	}
}
