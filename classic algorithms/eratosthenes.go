package main

import "fmt"

func main() {
	fmt.Println("enter the number of primes to waht number?")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		panic(err)
	}

	prime := make([]bool, num+1)

	for i := 2; i*i <= num; i++ {
		if prime[i] == false {
			for j := i * 2; j <= num; j += i {
				prime[j] = true // cross
			}
		}

	}

	for i := 2; i <= num; i++ {
		if prime[i] == false {
			fmt.Println(i)
		}
	}
}
