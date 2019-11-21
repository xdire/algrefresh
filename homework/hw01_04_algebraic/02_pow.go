package hw01_04_algebraic

import "fmt"

/*
	Power through iteration
	---
 */
func PowIter(a float64, b uint32) float64 {
	res := float64(1)
	for i:=uint32(0); i < b; i++ {
		res *= a
	}
	return res
}

/*
	Power calculated through quadratic multiplication
	with addition
	---
*/
func Pow2Mul(a float64, b int32) float64  {
	res := a
	pow := int32(1)
	// Until the power reaches half of desired power
	// double it, if reaches the half, then doubling
	// will not be possible anymore
	for pow < b / 2 {
		res *= res
		pow *= 2
		fmt.Printf("\nPow 1/2: %d Result: %f", pow, res)
	}
	// Increase the rest in iterative mode
	for pow < b {
		res *= a
		pow++
		fmt.Printf("\nPow 2/2: %d Result: %f", pow, res)
	}
	return res
}

/*
	Power calculated through counting power in
	binary disassemble fashion
	---
*/
func Pow2Bin(a float64, b int64) float64 {
	res := float64(1)
	// Check until b reaches 0 or 1
	for b > 1 {
		// If b is odd then multiply the result into the accumulator
		// - basically apply the single exponent each time the 1 is leftover of power
		if b % 2 == 1 {
			res *= a
		}
		// Power the value quadratically each cycle, meaning every
		// half-decreased power will be accounted and applied
		a *= a
		// Decrease the power in half
		b /= 2
	}
	// Fulfill the rest accumulated values into result
	if b > 0 {
		res *= a
	}
	return res
}
