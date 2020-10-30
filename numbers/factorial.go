package main

import "fmt"

func main() {
	fmt.Println("enter a number")

	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		panic(err)
	}

	fmt.Println(getFactorialRecursively(number))
}

func getFactorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}

	fac := number
	for number > 1 {
		number--
		fac = fac * number
	}
	return fac
}

func getFactorialRecursively(number int) int {
	if number == 0 {
		return 1
	}

	return number * getFactorialRecursively(number-1)
}
