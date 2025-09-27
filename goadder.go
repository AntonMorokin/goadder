// Package goadder provides simple integer addition functionality.
package goadder

// Add returns the sum of two integers.
//
// Parameters:
//   a - the first integer to add
//   b - the second integer to add
//
// Returns:
//   The sum of a and b
//
// Example:
//   sum := Add(5, 3) // sum == 8
//   sum := Add(-2, 7) // sum == 5
// See [funny link]: https://mathsisfun.com/numbers/addition.html
func Add(a, b int) int {
	return a + b
}
