package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	defer fmt.Println("world1")
	defer fmt.Println("world2")
	defer fmt.Println("world3")
	mydefer()
} // defer works on lifo - last in first out
// o/p:  hello
// world3
// world2
// world1
// world

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Hello babes...", i)
	}
}

// o/p: hello
// Hello babes... 4
// Hello babes... 3
// Hello babes... 2
// Hello babes... 1
// Hello babes... 0
// world3
// world2
// world1
// world
