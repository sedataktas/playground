package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := "abcdefghijklmnopqrstuvwxyz"
	b := "abcdefghijklmnopqrstuvwxyz"
	str := TwoToOne(a, b)
	fmt.Println(str)
}

func TwoToOne(s1 string, s2 string) string {
	chars := strings.Split(s1+s2, "")
	sort.Strings(chars)
	result := ""
	for _, s := range chars {
		chr := string(s)
		if !strings.Contains(result, chr) {
			result = result + chr
		}
	}
	return result
}
