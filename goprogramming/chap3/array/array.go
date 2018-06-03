package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hi"
	a[1] = "Man"
	fmt.Println(a[0], a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
