package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}

func hello() {
	if true {
		fmt.Println("Hello, World!")

		switch {
		case true:
			fmt.Println("Hello, World!")
		case false:
			fmt.Println("Hello, World!")
		}
	}
}
