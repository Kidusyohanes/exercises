package main

import "fmt"

func factorial(num int) int {
	var i, j int
	for i = 1; i < num; i++ {
		j = j * i
	}
	return j
}

func main() {
	num := 2
	j := factorial(num)
	fmt.Printf("The factorial of %d is %d\n", num, j)
}
