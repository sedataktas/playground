package main

import (
	"fmt"
)

var ones = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var tens = map[int]string{
	10: "ten",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

var others = map[int]string{
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
}

func main() {
	fmt.Println("enter a number")
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		panic(err)
	}

	if number < 0 || number >= 1000000 {
		panic("number should be between 0 and 1000000")
	}

	fmt.Println(num2word(number))
}

func num2word(num int) string {
	if num >= 0 && num < 10 {
		return ones[num]
	}

	if num >= 10 && num < 100 {
		return getTens(num)
	}

	if num >= 100 && num < 1000 {
		return getHundreds(num)
	}

	if num >= 1000 && num < 10000 {
		thousand := num / 1000
		thousandRemain := num % 1000

		word := ones[thousand] + " thousand "
		if thousandRemain > 0 {
			word += getHundreds(thousandRemain)

		}
		return word
	}

	if num >= 10000 && num < 100000 {
		thousand := num / 1000
		thousandRemain := num % 1000

		word := getTens(thousand) + " thousand "
		if thousandRemain > 0 {
			word += getHundreds(thousandRemain)
		}
		return word
	}

	if num >= 100000 && num < 1000000 {
		thousand := num / 1000
		thousandRemain := num % 1000

		word := getHundreds(thousand) + " thousand "
		if thousandRemain > 0 {
			word += getHundreds(thousandRemain)
		}
		return word
	}

	return ""
}

func getTens(num int) string {
	word := ""
	if num >= 10 && num < 20 {
		word = others[num]
		if word == "" {
			word = tens[num]
		}
		return word
	} else {
		ten := num / 10
		one := num % 10
		word = tens[ten*10]

		if one > 0 {
			word += "-" + ones[one]
		}
	}

	return word
}

func getHundreds(num int) string {
	hundred := num / 100
	ten := num % 100

	word := ones[hundred] + " hundred "
	if ten > 0 {
		word += getTens(ten)
	}
	return word
}
