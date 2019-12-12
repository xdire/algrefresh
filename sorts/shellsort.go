package sorts

import "fmt"

func ShellInt(a []int64)  {
	size := len(a)
	// Start step loop decreasing the steps in half
	for step := size / 2; step > 0; step = step / 2 {
		// Walk up all elements starting the steps start
		for i := step; i < size; i++ {
			// Save current result
			stor := a[i]
			j := i
			// While j is after beginning and (a[j - Step value]) greater than
			// current a[i] value do the swap
			for ; j >= step && a[j - step] > stor; j -= step {
				fmt.Printf("\n [j] going as %d with val [%d] with step %d with stor %d", j, a[j], step, stor)
				a[j] = a[j - step]
			}
			fmt.Printf("\n Placing a[j] going as %d with step %d with stor %d", j, step, stor)
			// Apply previous value to j-index
			a[j] = stor
		}
	}
}