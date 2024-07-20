package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 par")
	} else {
		fmt.Println("7 impar")
	}

	if 8%4 == 0 {
		fmt.Println("8 divisivel por 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 ou 7 divisivel por 2")
	}

	if num := -10; num < 0 {
		fmt.Println("Number is negative")
	} else if num < 10 {
		fmt.Println("Number has 1 digit")
	} else {
		fmt.Println("Number has multiple digits")
	}
}
