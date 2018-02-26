package roulette


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
}
