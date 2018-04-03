package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Print(adder()(1))
	c := adder();
	print(c(1));
	print(c(1));
	print(c(1));
}
