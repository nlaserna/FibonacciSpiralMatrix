package helper

import "fmt"

func FibonacciSequence(n int) (sequence []int) {
	t1 := 0
	t2 := 1
	nextTerm := 0

	sequence = make([]int, n)
	fmt.Print("Fibonacci Series :")
	for i := 1; i <= n; i++ {
		if i == 1 {
			sequence[i-1] = t1
			fmt.Print(" ", sequence[i-1])
			continue
		}
		if i == 2 {
			sequence[i-1] = t2
			fmt.Print(" ", sequence[i-1])
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
		sequence[i-1] = nextTerm
		fmt.Print(" ", sequence[i-1])
	}
	fmt.Println("")
	return sequence
}
