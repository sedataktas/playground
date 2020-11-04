package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter the string you want reversed")
	var s string
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		panic(err)
	}

	strArr := strings.Split(s, "")

	var reversedArr []string
	for i := len(strArr) - 1; i >= 0; i-- {
		reversedArr = append(reversedArr, strArr[i])
	}

	reversed := strings.Join(reversedArr, "")
	fmt.Println(reversed)
}
