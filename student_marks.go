package main

import "fmt"

func main() {

	var marks int32
	fmt.Println("enter student marks")
	fmt.Scanf("%d", &marks)
	if marks > 95 {
		fmt.Println("Grade A")

	} else if marks > 75 {
		fmt.Println("Grade B")

	} else if marks > 65 {
		fmt.Println("Grade C")

	} else {
		fmt.Println("fail")
	}

}
