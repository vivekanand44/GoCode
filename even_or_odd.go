package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println("The number is Even =", i)

		} else {
			fmt.Println("The number is Odd =", i)
		}
	}
}
