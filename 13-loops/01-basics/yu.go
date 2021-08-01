package main

import "fmt"

func yu_main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("The sum is %v", sum)
}
