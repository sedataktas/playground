package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("how many times would you like to flip the coin")
	var count int
	_, err := fmt.Scanf("%d", &count)
	if err != nil {
		panic(err)
	}

	tail := 0
	head := 1

	var tails []int
	var heads []int

	for count > 0 {
		random := rand.Intn(2)
		if random == tail {
			tails = append(tails, random)
		}

		if random == head {
			heads = append(heads, random)
		}
		count--
	}

	fmt.Printf("Tails : %d , heads : %d\n", len(tails), len(heads))
}
