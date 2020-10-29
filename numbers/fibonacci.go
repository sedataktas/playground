package main

import (
	"fmt"
)

const MaxFibDigit = 78

func main() {
	fmt.Println("Enter Nth Digit:")

	var digit int
	_, err := fmt.Scanf("%d", &digit)
	if err != nil {
		panic(err)
	}

	if digit > 0 && digit < MaxFibDigit {
		fmt.Println(getFib(digit))
	} else {
		fmt.Printf("Digit number should be between 1 and %d\n", MaxFibDigit)
	}
}

func getFib(digit int) []int {
	firstNum := 0
	secondNum := 1
	var fib []int
	if digit == 1 {
		fib = append(fib, firstNum)
		return fib
	}
	if digit == 2 {
		fib = append(fib, firstNum, secondNum)
		return fib
	}

	fib = append(fib, firstNum, secondNum)
	for i := len(fib); i < digit; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib
}
