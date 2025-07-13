package main

import "fmt"

func main() {
	fmt.Println("first way if else condition")

	age := 28

	if age > 18 {
		fmt.Println("red team member")
	} else if age < 18 {
		fmt.Println("blue team member")
	} else {
		fmt.Println("pink team member")
	}

	fmt.Println("Second way if else condition")
	if 15%2 == 0 {
		fmt.Println("even number")
	} else {
		fmt.Println("odd number")
	}

	fmt.Println("third way if else condition")
	if num := 15; num < 15 {
		fmt.Println("eat burger")
	} else {
		fmt.Println("dosa")
	}
}
