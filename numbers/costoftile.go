package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the floor length")
	var floorLength int
	_, err := fmt.Scanf("%d", &floorLength)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter the floor width")
	var floorWidth int
	_, err = fmt.Scanf("%d", &floorWidth)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter the length of edge of the tile")
	var tileLength int
	_, err = fmt.Scanf("%d", &tileLength)
	if err != nil {
		panic(err)
	}

	if tileLength > floorWidth || tileLength > floorLength {
		panic("tile length can not be bigger than floor length or width")
	}

	fmt.Println("Enter the price of one tile")
	var tilePrice int
	_, err = fmt.Scanf("%d", &tilePrice)
	if err != nil {
		panic(err)
	}

	floorArea := floorLength * floorWidth
	tileArea := tileLength * tileLength
	cost := (floorArea / tileArea) * tilePrice

	fmt.Printf("Total cost : %d\n", cost)
}
