package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
			continue
		}
		if i%5 == 0 && i%3 != 0 {
			fmt.Println("buzz")
			continue
		}
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		fmt.Println(i)
	}
}
