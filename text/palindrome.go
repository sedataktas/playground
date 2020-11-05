package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter the word")
	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		panic(err)
	}

	strArr := strings.Split(word, "")

	var reversedArr []string
	for i := len(strArr) - 1; i >= 0; i-- {
		reversedArr = append(reversedArr, strArr[i])
	}

	reversed := strings.Join(reversedArr, "")
	if word == reversed {
		fmt.Printf("entered word :%s is palindrome\n", word)
	} else {
		fmt.Printf("entered word :%s is not palindrome\n", word)
	}
}
