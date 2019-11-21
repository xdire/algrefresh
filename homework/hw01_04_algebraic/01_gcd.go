package hw01_04_algebraic

/*
	GCD by subtraction
	---
	Example of 32 56
	32 != 56
	  -> 56 - 32
		 -> 32 - 24
			-> 24 - 8
			   -> 16 - 8 -> 8 = 8
 */
func GCDSub(a int64, b int64) int64 {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

/*
	GCD by modulo
	---
	Example of 32 56
	32 != 0 && 56 != 0
	  -> 56 % 32 = 24
		 -> 32 % 24 = 8
			-> 24 % 8 = 0
			   -> a > 0 ? return a(8)
*/
func GCDMod(a int64, b int64) int64   {
	for a != 0 && b != 0 {
		if a > b {
			a %= b
		} else {
			b %= a
		}
	}
	if a != 0 {
		return a
	}
	return b
}
