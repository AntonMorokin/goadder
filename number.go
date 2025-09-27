package goadder

import (
	"golang.org/x/exp/constraints"
)

// Number is a type constraint that represents all numeric types.
// It includes both integer types (int, int8, int16, int32, int64, uint, etc.)
// and floating-point types (float32, float64).
type Number interface {
	constraints.Integer | constraints.Float
}
