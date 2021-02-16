package main

import "fmt"

func main() {
	fmt.Println("final :" + comma("12345678900"))
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	fmt.Println(s[:n-3])
	return comma(s[:n-3]) + "," + s[n-3:]
}
