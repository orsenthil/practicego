// [_Variadic functions_](https://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

// Here's a function that will take an arbitrary number
// of `int`s as arguments.

// TODO: Create function sum(nums ...int) that calculates sum of all nums
// Within the function, the type of `nums` is equivalent to `[]int`. We can call `len(nums)`,
// We can iterate over it with `range`, etc.


func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.

	// TODO: Call sum(1, 2) and sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.

	// TODO: Create slice nums := []int{1, 2, 3, 4}
	// TODO: Call sum(nums...)

}