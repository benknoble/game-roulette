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
}
