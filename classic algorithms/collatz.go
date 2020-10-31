package main

import "fmt"

func main() {
	fmt.Println("enter the number that you want to reduce to 1")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}

	var stepsCount int
	for {
		if num == 1 {
			break
		}
		if isEven(num) {
			num = num / 2
			fmt.Println(num)
		} else {
			num = num*3 + 1
			fmt.Println(num)
		}
		stepsCount++
	}
	fmt.Printf("total steps : %d\n", stepsCount)
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}
