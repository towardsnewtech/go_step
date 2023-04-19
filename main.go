package main

import (
	"fmt"
)

var count int = 0

func step(total int, sum int) int {
	if sum > total {
		return 0
	} else if sum == total {
		count ++
	} else {
		step(total, sum+1)
		step(total, sum+2)
	}
	return 1
}

func main() {
	step(4, 0)
	fmt.Printf("methods are: %d\n", count)
}