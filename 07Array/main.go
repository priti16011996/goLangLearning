package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array")
	var fruitList [4]string
	fruitList[0] = "Mango"
	fruitList[1] = "banana"
	fruitList[3] = "Apple"
	fmt.Println("Fruitlist arraycontains: ", fruitList)

	fmt.Println("Second way to declare array:")

	var MenuItems = [5]string{"chole bhature", "Biryani", "Chuda Matar", "Sahi Paner", "gulab Jamun"}
	fmt.Println("Menu items: ", MenuItems)

}
