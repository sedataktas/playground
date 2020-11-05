package main

import (
	"fmt"
	"strings"
)

var vowels = [5]string{"a", "e", "i", "o", "u"}

func main() {
	fmt.Println("enter the word that find the vowels")
	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		panic(err)
	}

	strArr := strings.Split(word, "")

	var vowelsInWord []string
	for _, s := range strArr {
		for _, v := range vowels {
			if s == v {
				vowelsInWord = append(vowelsInWord, v)
			}
		}
	}

	fmt.Println(vowelsInWord)
}
