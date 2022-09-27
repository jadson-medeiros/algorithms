package main

import "fmt"

// The sequence being 1 1 2 3 5 8, if value=6, the answer is 8

func main() {

	fmt.Println("The value iterative is:", fibonacciIt(6))

	fmt.Println("The value recursive is:", fibonacciRec(6))
}

func fibonacciIt(n int) int {

	count := 1

	previous := 0
	current := 1
	next := 0

	for count < n {
		next = current + previous

		previous = current
		current = next
		count++
	}

	return next
}

func fibonacciRec(n int) int {

	if n <= 1 {
		return n
	}

	value := fibonacciRec(n-1) + fibonacciRec(n-2)

	return value
}
