// Package libgollatz will take a n integer, perform the collatz aalgorithm on it,
// and then put each value ithits into a linked list, which will then be
// put into an array. Then it will return a Result consisting of the original
// value and the array of values.
package libgollatz

import "container/list"

// Result is made up of the initial value and an array containing all
// of the steps to get to 1 using the Collatz algorithm
type Result struct {
	Value uint64
	Steps []uint64
}

// Collatz will take in a number, and then run it through the Collatz
// algorithm (3n+1). It returns a Result which stores the initial value
// and the steps it took to get to 1.
func Collatz(start uint64) Result {
	var result Result
	num := start
	stepsTaken := list.New()
	result.Value = start

	for num > 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		stepsTaken.PushBack(num)
	}

	// Initialize a byte slice to hold the steps
	steps := make([]uint64, stepsTaken.Len())

	// Holds the current place in the list so that we can add them to
	count := 0

	// Put the values from our list of steps into an array.
	for e := stepsTaken.Front(); e != nil; e = e.Next() {
		steps[count] = e.Value.(uint64)
		count++
	}

	result.Steps = steps

	return result
}
