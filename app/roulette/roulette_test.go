package roulette

import (
	"math"
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

func TestRandAsFloater(t *testing.T) {
	var f Floater = rand.New(rand.NewSource(1))
	c := FairCoin{}
	c.GenerateRandom(f)
}

func TestFairCoin(t *testing.T) {
	c := FairCoin{}
	var f Floater = newMockFloater()
	for i := 0; i < 10; i++ {
		heads := c.GenerateRandom(f)
		var result int
		if i < 5 {
			result = Heads
		} else {
			result = Tails
		}
		if heads != result {
			t.Errorf("Expected %v, got %v", result, heads)
		}
	}
}

func TestBiasedCoin(t *testing.T) {
	c := NewBiasedCoin(.3)
	var f Floater = newMockFloater()
	for i := 0; i < 10; i++ {
		heads := c.GenerateRandom(f)
		var result int
		if i < 3 {
			result = Heads
		} else {
			result = Tails
		}
		if heads != result {
			t.Errorf("Expected %v, got %v", result, heads)
		}
	}
}

func TestFairDie(t *testing.T) {
	d := NewFairDie(10)
	var f Floater = newMockFloater()
	for i := 0; i < 10; i++ {
		face := d.GenerateRandom(f)
		result := i + 1
		if face != result {
			t.Errorf("Expected %v, got %v", result, face)
		}
	}
}

func TestLoadedDie(t *testing.T) {
	_, err := NewLoadedDie(4, []float64{0.1, 0.3, 0.6})
	if err == nil {
		t.Error("Expected error with missing dies, got nil")
	}
	_, err = NewLoadedDie(3, []float64{0.1, 0.3, 0.6})
	if err != nil {
		t.Errorf("Got error on creation: %v", err)
	}
}

func TestVoseInit(t *testing.T) {
	d, err := NewLoadedDie(3, []float64{0.1, 0.3, 0.6})
	if err != nil {
		t.Errorf("Got error on creation: %v", err)
	}
	alias, prob := voseInit(d)
	expectedAlias := []int{2, 2, 0}
	expectedProb := []float64{0.3, 0.9, 1}
	if !equals(alias, expectedAlias) && !equalsFloat64s(prob, expectedProb) {
		t.Errorf("Expected (%v,%v), got (%v,%v)", expectedAlias, expectedProb, alias, prob)
	}

	d, err = NewLoadedDie(7, []float64{1.0 / 8.0, 1.0 / 5.0, 1.0 / 10.0, 1.0 / 4.0, 1.0 / 10.0, 1.0 / 10.0, 1.0 / 8.0})
	if err != nil {
		t.Errorf("Got error on creation: %v", err)
	}
	alias, prob = voseInit(d)
	expectedAlias = []int{1, 0, 3, 1, 3, 3, 3}
	expectedProb = []float64{7.0 / 8.0, 1.0, 7.0 / 10.0, 29.0 / 40.0, 7.0 / 10.0, 7.0 / 10.0, 7.0 / 8.0}
	if !equals(alias, expectedAlias) && !equalsFloat64s(prob, expectedProb) {
		t.Errorf("Expected (%v,%v), got (%v,%v)", expectedAlias, expectedProb, alias, prob)
	}
}

func TestPop(t *testing.T) {
	stack := []int{1, 2, 3}
	head, tail := pop(stack)
	expHead, expTail := 3, []int{1, 2}
	if head != expHead || !equals(tail, expTail) {
		t.Errorf("Expected (%v,%v), got (%v,%v)", expHead, expTail, head, tail)
	}
}

func TestLoadedDieGenerateRandom(t *testing.T) {
	t.Errorf("Not yet implemented")
}

func equals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalsFloat64s(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > math.SmallestNonzeroFloat64 {
			return false
		}
	}
	return true
}
