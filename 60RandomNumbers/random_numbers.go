// Go's `math/rand/v2` package provides
// [pseudorandom number](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// generation.

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// For example, `rand.IntN` returns a random `int` n,
	// `0 <= n < 100`.

	fmt.Println(rand.IntN(100), ",")
	fmt.Println(rand.IntN(100))

	// `rand.Float64` returns a `float64` `f`,
	// `0.0 <= f < 1.0`.

	fmt.Println(rand.Float64())

	// This can be used to generate random floats in
	// other ranges, for example `5.0 <= f' < 10.0`.

	fmt.Println((rand.Float64()*5)+5, ",")
	fmt.Println((rand.Float64() * 5) + 5)

	// If you want a known seed, create a new
	// `rand.Source` and pass it into the `New`
	// constructor. `NewPCG` creates a new
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// source that requires a seed of two `uint64`
	// numbers.

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Println(r2.IntN(100), ",")
	fmt.Println(r2.IntN(100))

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Println(r3.IntN(100), ",")
	fmt.Println(r3.IntN(100))
}