// Package goadder provides simple integer addition functionality.
package goadder

// Add returns the sum of two numbers of the same numeric type.
//
// The function supports all numeric types defined by the Number constraint.
// It provides compile-time type safety while allowing generic numeric operations.
//
// Example:
//
//	sum := Add(5, 3)        // returns 8 (int)
//	sum := Add(2.5, 3.7)    // returns 6.2 (float64)
//	sum := Add(int32(10), int32(20)) // returns 30 (int32)
//
// Parameters:
//   a - The first operand
//   b - The second operand
//
// Returns:
//   The sum of a and b, with the same type as the operands.
// See [funny link]
//
// [funny link]: https://mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
