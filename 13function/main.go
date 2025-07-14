package main

import "fmt"

func main() {
	fmt.Println("function tutorial")
	greet()
	result := adder(15, 85)
	fmt.Println("sun result:", result)
	result2 := proAdder(1, 2, 3, 4, 5)
	fmt.Println("proAdder result:", result2)

}
func greet() {
	fmt.Println("Hello,user")
}

func adder(a int, b int) int {
	sum := a + b
	return sum
}

func proAdder(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum

}
