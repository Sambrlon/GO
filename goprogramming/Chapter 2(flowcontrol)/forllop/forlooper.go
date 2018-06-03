package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	sum1 := 0
	//sum2 := 0

	start := time.Now()
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println(sum)

	for i := 0; i < 10000000000; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	//for i := 0; i < 10000000000000; i++ {
	//	sum2 += i
	//}
	//fmt.Println(sum2)
	fmt.Println(time.Since(start))
}
