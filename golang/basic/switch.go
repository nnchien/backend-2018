package main

import "fmt"

func main() {

	var (
		even  int
		odd   int
		total int
	)

	numbers := []int{1, 2, 3, 5, 0, 7, 8, 9, 10}

	for _, n := range numbers {
		total += 1

		switch {
		case n == 0:
		case n%2 == 0:
			even += 1
		default:
			odd += 1
		}
	}

	fmt.Printf("Even %d, odd %d, total %d\n", even, odd, total)
}
