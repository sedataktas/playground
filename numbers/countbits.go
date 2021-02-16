package main

import "fmt"

func main() {
	fmt.Println(CountBits(1234))
}

func CountBits(in uint) int {
	if in <= 0 {
		return 0
	}

	count := 0
	for in > 0 {
		if in%2 == 0 {
			in = in / 2
		} else {
			count++
			in = in / 2
		}
	}

	return count
}
