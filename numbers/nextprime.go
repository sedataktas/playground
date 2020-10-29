package main

import (
	"bufio"
	"fmt"
	"os"
)

const firstPrimeNumber = 2

func main() {
	// Firstly, we write first prime number
	fmt.Printf("First prime number is %d\n", firstPrimeNumber)
	prevPrimeNumber := firstPrimeNumber

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Do you want to show next prime number?(y/n)")
	for scanner.Scan() {
		switch scanner.Text() {
		case "y":
			nextPrimeNumber := getNextPrimeNumber(prevPrimeNumber)
			fmt.Printf("Next prime number:%d\n", nextPrimeNumber)
			prevPrimeNumber = nextPrimeNumber
		case "n":
			os.Exit(0)
		default:
			fmt.Println("entered wrong input")
		}
		fmt.Println("Do you want to show next prime number?(y/n)")
	}
	if scanner.Err() != nil {
		panic("an error occurred when read input")
	}
}

func getNextPrimeNumber(prevPrimeNumber int) int {
	for {
		prevPrimeNumber++
		countOfNotDivide := 0
		i := 2
		for i < prevPrimeNumber {
			if prevPrimeNumber%i == 0 {
				break
			}
			if prevPrimeNumber%i != 0 {
				countOfNotDivide++
			}
			i++
		}
		if prevPrimeNumber-2 == countOfNotDivide {
			break
		}
	}
	return prevPrimeNumber
}
