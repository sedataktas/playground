package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter decimal number")

	var decNum int
	_, err := fmt.Scanf("%d", &decNum)
	if err != nil {
		panic(err)
	}
	fmt.Printf("converted to binary:%v\n", decimalToBinary(decNum))

	fmt.Println("Enter binary number")
	var binNum int
	_, err = fmt.Scanf("%d", &binNum)
	if err != nil {
		panic(err)
	}
	fmt.Printf("converted to decimal:%d\n", binaryToDecimal(binNum))
}

func decimalToBinary(num int) []int {
	var reverseBinaryNums []int
	for {
		division := num / 2
		remain := num % 2
		reverseBinaryNums = append(reverseBinaryNums, remain)
		num = division
		if division == 1 || division == 0 {
			reverseBinaryNums = append(reverseBinaryNums, division)
			break
		}
	}

	var binaryNums []int
	for i := len(reverseBinaryNums) - 1; i >= 0; i-- {
		binaryNums = append(binaryNums, reverseBinaryNums[i])
	}
	return binaryNums
}

func binaryToDecimal(num int) int {
	strBinaryNum := strconv.Itoa(num)
	strArrBinaryNum := strings.Split(strBinaryNum, "")

	var decimalNumber float64
	pow := 0.0
	for i := len(strArrBinaryNum) - 1; i >= 0; i-- {
		intNum, err := strconv.Atoi(strArrBinaryNum[i])
		if err != nil {
			panic(err)
		}
		decimalNumber += float64(intNum) * math.Pow(2, pow)
		pow++
	}
	return int(decimalNumber)
}
