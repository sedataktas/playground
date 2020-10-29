package main

import "fmt"

func main() {
	fmt.Println("Enter the number that will find prime factorization:")

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}

	primes := findPrimes(num)
	if len(primes) == 0 {
		fmt.Println("there is no prime factors")
	} else {
		fmt.Println(getPrimeFactors(primes, num))
	}
}

func findPrimes(num int) []int {
	var primes []int
	for i := 2; i < num; i++ {
		if num%i == 0 {
			primes = append(primes, i)
		}
	}
	return primes
}

func getPrimeFactors(primes []int, num int) []int {
	var factors []int
	for _, p := range primes {
		for num%p == 0 {
			factors = append(factors, p)
			num = num / p
		}
	}
	return factors
}
