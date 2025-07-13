package main

import "fmt"

func main() {
	fmt.Println("Learn Pointer in address")
	// var pnt *int
	// fmt.Println(pnt)

	var num int = 1996
	var pnt *int = &num
	fmt.Println("Print what strore in pnt: ", pnt)
	fmt.Println("Print what strore in *pnt: ", *pnt)

	*pnt = *pnt + 10
	fmt.Println("Print what strore in pnt: ", num)
}
