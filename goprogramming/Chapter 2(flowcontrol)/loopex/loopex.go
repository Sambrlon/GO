package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	counter := 0
	start := 1.0
	finish := math.Sqrt(x)
	for math.Abs(finish-start) > 0.00000000000001 {
		start = start - (start*start-x)/(2*start)
		counter = counter + 1

	}
	fmt.Println(math.Abs(finish - start))
	fmt.Printf("Procss finished with %d iterations \n", counter)
	fmt.Println(counter)
	return start
}

func main() {
	fmt.Println(math.Sqrt(1000157), sqrt(1000157))
}
