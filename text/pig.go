package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("enter the word that convert pig latin")
	var word string
	_, err := fmt.Scanf("%s", &word)
	if err != nil {
		panic(err)
	}

	strArr := strings.Split(word, "")
	var pig []string
	for i := 1; i < len(strArr); i++ {
		pig = append(pig, strArr[i])
	}

	pig = append(pig, "-", strArr[0], "ay")
	fmt.Println(strings.Join(pig, ""))
}
