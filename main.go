package main

import (
	"fmt"
)

/*
func main() {

	fmt.Println("HELLO")
	a, b := 13.0, 3.0
	fmt.Println(calc.Add(a, b))
	fmt.Println(calc.Sub(a, b))
	fmt.Println(calc.Mul(a, b))
	fmt.Println(calc.Div(a, b))

}
*/
func Prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	var primes []int

	for i := 1; i <= 10; i++ {
		if Prime(i) {
			primes = append(primes, i)
			sum += i
		}
	}

	fmt.Printf("The prime numbers between 1 and 10 are: %v\n", primes)
	fmt.Printf("The sum of prime numbers between 1 and 10 is: %d\n", sum)
}
