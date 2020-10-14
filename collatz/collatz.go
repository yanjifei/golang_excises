//Start with a number n > 1.
//Find the number of steps it takes to reach one using the following process:
//If n is even, divide it by 2. If n is odd, multiply it by 3 and add 1.

package main

import "fmt"

func main() {
	n := 0
	step := 0

	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &n)

	_, result := collatz2(n, 0)
	fmt.Printf("Result = %d\r\n", result)

	for {
		if n <= 1 {
			fmt.Printf("Number of steps: %d\n", step)
			break
		} else {
			collatz(&n, &step)
		}
	}

}

func collatz(n, step *int) {
	if (*n)%2 == 0 {
		(*n) /= 2
	} else {
		(*n) = (*n)*3 + 1
	}
	(*step)++
}

func collatz2(n, step int) (int, int) {
	a := n % 2
	if n <= 1 {
		return n, step
	}
	switch a {
	case 0:
		return collatz2(n/2, step+1)
	case 1:
		return collatz2(n*3+1, step+1)
	default:
		return n, step
	}
}
