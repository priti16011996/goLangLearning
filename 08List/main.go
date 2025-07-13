package main

import "fmt"

func main() {
	fmt.Println("Welcome to List Tut")
	var list = []string{"Mix Veg", "Sahi Paneer", "Pulao", "Papad"}
	fmt.Printf("list %T\n", list)
	list = append(list, "Burger", "Fries")
	fmt.Println("list %T\n", list)
	list = append(list[1:3])
	fmt.Println("list %T\n", list)
}
