package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func SqrtAgain(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for n := 0; n < 10; n++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func Errors() {
	fmt.Println(SqrtAgain(2))
	fmt.Println(SqrtAgain(-2))
}
