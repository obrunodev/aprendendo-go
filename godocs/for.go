package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("Simple", i)
		i++
	}
	for j := 0; j < 10; j++ {
		fmt.Println("Initial/condition/after", j)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("for break")
		break
	}
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("With condition", n)
	}
}
