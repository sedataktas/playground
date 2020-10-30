package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(getFirst8HappyNumbers())
}

func getFirst8HappyNumbers() []int {
	number := 1
	var happyNumbers []int

	for len(happyNumbers) < 8 {
		if number == 1 {
			happyNumbers = append(happyNumbers, number)
		}

		strArrNum := intToStringAndSplit(number)
		for len(strArrNum) > 1 {
			total := 0
			for _, num := range strArrNum {
				total = calculatePow(total, num)
			}

			if total == 1 {
				happyNumbers = append(happyNumbers, number)
				strArrNum = []string{}
			} else {
				strArrNum = intToStringAndSplit(total)
			}

		}
		number++
	}

	return happyNumbers
}

func intToStringAndSplit(number int) []string {
	strNum := strconv.Itoa(number)
	return strings.Split(strNum, "")
}

func calculatePow(total int, num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	total = total + n*n
	return total
}
