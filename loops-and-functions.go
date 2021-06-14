package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var previous float64
	z := x
	for n := 0; n < 10; n++ {
		previous = z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Step %v: z = %v\n", n+1, z)
		if previous == z {
			fmt.Println("Stopping early")
			return z
		}
	}
	return z
}

func LoopsAndFunctions() {
	fmt.Println(Sqrt(40))
	fmt.Println(math.Sqrt(40))
}
