package roulette

import (
	"errors"
	"math"
)

// A Floater can provide a float64 on each invocation of it's method
type Floater interface {
	Float64() float64
}

// A RandomGenerator generates a random number using the results of Floater's Float64 method
type RandomGenerator interface {
	// Generate a random int based on the Floater's random data and the underlying type
	GenerateRandom(Floater) int
}

// ErrorSides is a value representing a lack of the necessary probabilities for
// a loaded die
var ErrorSides = errors.New("Not enough probabilities provided for the number of sides")

// Heads value of a coin
var Heads = 1

// Tails value of a coin
var Tails int

// A FairCoin gives 1 or 0 as a RandomGenerator
type FairCoin struct{}

// GenerateRandom a Heads or a Tails, like flipping a coin
func (c FairCoin) GenerateRandom(f Floater) int {
	x := f.Float64()
	if x < 0.5 {
		return Heads
	} // x >= 0.5
	return Tails
}

// A BiasedCoin is like a FairCoin, only heads and tails have (possibly)
// differing probabilities of showing
type BiasedCoin struct {
	headsProbability float64
}

// NewBiasedCoin creates a BiasedCoin with probability `headsProbability` of a
// heads showing
func NewBiasedCoin(headsProbability float64) BiasedCoin {
	return BiasedCoin{headsProbability}
}

// GenerateRandom a Heads or a Tails, like flipping a coin, only this flip is
// biased after the coin
func (c BiasedCoin) GenerateRandom(f Floater) int {
	x := f.Float64()
	if x < c.headsProbability {
		return Heads
	} // x >= c.headsProbability
	return Tails
}

// A FairDie gives an integer in the range [1,number of sides] as a
// RandomGenerator
type FairDie struct {
	sides int
}

// NewFairDie creates a new FairDie with n sides
func NewFairDie(n int) FairDie {
	return FairDie{n}
}

// GenerateRandom an integer in the range [1,number of sides on the die]
func (d FairDie) GenerateRandom(f Floater) int {
	x := f.Float64()
	return int(math.Floor(x * float64(d.sides)))
}

// A LoadedDie gives an integer in the range [1,number of sides] as a
// RandomGenerator, but respects the probability given to each side
type LoadedDie struct {
	sides         int
	probabilities []float64
}

// NewLoadedDie creates a new LoadedDie with n sides, where side i has
// probability ps[i] of showing
func NewLoadedDie(n int, ps []float64) (LoadedDie, error) {
	if len(ps) != n {
		return LoadedDie{}, ErrorSides
	}
	return LoadedDie{n, ps}, nil
}

// GenerateRandom an integer in the range [1,number of sides] respecting the
// weight's of the dice
//
// Implented based on the alias method described here:
// http://www.keithschwarz.com/darts-dice-coins/
func (d LoadedDie) GenerateRandom(f Floater) int {
	alias, prob := voseInit(d)
	i := NewFairDie(d.sides).GenerateRandom(f)
	flip := NewBiasedCoin(prob[i]).GenerateRandom(f)
	if flip == Heads {
		return i
	} //flip == Tails
	return alias[i]
}

func voseInit(d LoadedDie) (alias []int, prob []float64) {
	alias, prob = make([]int, d.sides), make([]float64, d.sides)
	var probabilities []float64
	copy(d.probabilities, probabilities)
	probabilities = d.probabilities[:]
	for i := range probabilities {
		probabilities[i] = probabilities[i] * float64(d.sides)
	}
	var small, large []int
	for i, p := range probabilities {
		scaledP := p
		if scaledP < 1 {
			small = append(small, i)
		} else {
			large = append(large, i)
		}
	}
	for len(small) != 0 && len(large) != 0 {
		var l, g int
		l, small = pop(small)
		g, large = pop(large)
		prob[l] = probabilities[l]
		alias[l] = g
		probabilities[g] = probabilities[g] + probabilities[l] - 1
		if probabilities[g] < 1 {
			small = append(small, g)
		} else {
			large = append(large, g)
		}
	}
	for len(large) != 0 {
		var g int
		g, large = pop(large)
		prob[g] = 1
	}
	for len(small) != 0 {
		var l int
		l, small = pop(small)
		prob[l] = 1
	}
	return alias, prob
}

func pop(s []int) (head int, tail []int) {
	len := len(s)
	if len != 0 {
		head = s[len-1]
		tail = s[:len-1]
	}
	return
}
