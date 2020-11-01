package main

import "fmt"

func main() {
	var numbers []int
	numbers = append(numbers, 3, 5, 10, 9, 12, 20)
	fmt.Println(bubbleSort(numbers))
}

func bubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1] < numbers[i] {
			numbers[i+1], numbers[i] = numbers[i], numbers[i+1]
		}
	}
	return numbers
}
