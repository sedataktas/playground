package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("How many seconds later would you like to set the alarm?")
	var sec int
	_, err := fmt.Scanf("%d", &sec)
	if err != nil {
		panic(err)
	}

	timeSec := rand.Int31n(int32(sec))
	timer1 := time.NewTimer(time.Duration(timeSec) * time.Second)
	<-timer1.C

	fmt.Println("\a")
}
