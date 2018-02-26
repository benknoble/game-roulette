package roulette

import (
	"math/rand"
	"testing"
)

type mockFloater struct {
	index int
	vals  []float64
}

func newMockFloater() *mockFloater {
	fs := []float64{
		0.0,
		0.1,
		0.2,
		0.3,
		0.4,
		0.5,
		0.6,
		0.7,
		0.8,
		0.9,
		1.0,
	}
	return &mockFloater{0, fs}
}

func (f *mockFloater) Float64() float64 {
	if f.index >= len(f.vals) {
		f.index = 0
	}
	defer func() { f.index++; return }()
	return f.vals[f.index]
}

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
			t.Errorf("New Random seeded with 1 gave (%v), expected (%v)", test, ten[i])
		}
	}
}
