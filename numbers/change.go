package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter cost")
	var cost float64
	_, err := fmt.Scanf("%f", &cost)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter payment")
	var payment float64
	_, err = fmt.Scanf("%f", &payment)
	if err != nil {
		panic(err)
	}

	if payment < cost {
		panic("payment can not be smaller than cost")
	}

	diff := payment - cost
	calculateAndPrint(diff)
}

func calculateAndPrint(diff float64) {
	fmt.Println("Change due:")

	twohundreds := math.Floor(diff / 200)
	leftover := diff - twohundreds*200
	if twohundreds > 0 {
		fmt.Println("$200 bills:", twohundreds)
	}

	hundreds := math.Floor(leftover / 100)
	leftover -= hundreds * 100
	if hundreds > 0 {
		fmt.Println("$100 bills:", hundreds)
	}

	fifties := math.Floor(leftover / 50)
	leftover -= fifties * 50
	if fifties > 0 {
		fmt.Println("$50 bills: ", fifties)
	}

	twenties := math.Floor(leftover / 20)
	leftover -= twenties * 20
	if twenties > 0 {
		fmt.Println("$20 bills: ", twenties)
	}

	tens := math.Floor(leftover / 10)
	leftover -= tens * 10
	if tens > 0 {
		fmt.Println("$10 bills: ", tens)
	}

	fives := math.Floor(leftover / 5)
	leftover -= fives * 5
	if fives > 0 {
		fmt.Println("$5 bills:  ", fives)
	}

	ones := math.Floor(leftover)
	leftover -= ones
	if ones > 0 {
		fmt.Println("$1 bills:  ", ones)
	}

	fiveOf1 := math.Floor(leftover / 0.5)
	leftover -= fiveOf1 * 0.5
	if fiveOf1 > 0 {
		fmt.Println("$0.5 bills:  ", fiveOf1)
	}

	twentyfiveOf1 := math.Floor(leftover / 0.25)
	leftover -= twentyfiveOf1 * 0.25
	if twentyfiveOf1 > 0 {
		fmt.Println("$0.25 bills:    ", twentyfiveOf1)
	}

	oneOf1 := math.Floor(leftover / 0.1)
	leftover -= oneOf1 * 0.1
	if oneOf1 > 0 {
		fmt.Println("$0.1 bills:   ", oneOf1)
	}

	fiveOfhundred := math.Floor(leftover / 0.05)
	leftover -= fiveOfhundred * 0.05
	if fiveOfhundred > 0 {
		fmt.Println("$0.05 bills:   ", fiveOfhundred)
	}

	oneOfHundred := math.Floor(leftover / 0.01)
	leftover -= oneOfHundred * 0.01
	if oneOfHundred > 0 {
		fmt.Println("$0.01 bills:   ", oneOfHundred)
	}
}
