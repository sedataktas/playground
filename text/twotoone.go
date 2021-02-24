package main

import (
	"fmt"
	"sort"
)

func main() {
	a := "abcdefghijklmnopqrstuvwxyz"
	b := "abcdefghijklmnopqrstuvwxyz"
	str := TwoToOne(a, b)
	fmt.Println(str)
}

func TwoToOne(s1 string, s2 string) string {
	var sum []byte
	sum = append(sum, s1[0])

	for _, s1v := range s1 {
		if !contains(sum, s1v) {
			sum = append(sum, byte(s1v))
		}
	}

	for _, s2v := range s2 {
		if !contains(sum, s2v) {
			sum = append(sum, byte(s2v))
		}
	}

	var strArr []string
	for _, s := range sum {
		strArr = append(strArr, string(s))
	}
	sort.Strings(strArr)

	var str string
	for _, s := range strArr {
		str += s
	}
	return str
}

func contains(s []byte, e int32) bool {
	for _, a := range s {
		if a == byte(e) {
			return true
		}
	}
	return false
}
