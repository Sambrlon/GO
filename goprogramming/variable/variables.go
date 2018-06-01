package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	var j, k = 1, 2
	fmt.Println(i, c, python, java)
	fmt.Println(j, i, k)
	var i1, j1 int = 1, 2
	k1 := 3
	c1, python1, java1 := true, false, "no!"

	fmt.Println(i1, j1, k1, c1, python1, java1)
}
