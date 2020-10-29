package main

import (
	"fmt"
	"math"
)

const (
	MaxDigit = 20
)

func main() {
	fmt.Println("Enter Nth Digit:")

	var digit int
	_, err := fmt.Scanf("%d", &digit)
	if err != nil {
		panic(err)
	}

	if digit > 0 && digit < MaxDigit {
		fmt.Printf("%.*f\n", digit, math.Pi)
	} else {
		fmt.Printf("Digit number should be between 0 and %d\n", MaxDigit)
	}
}
